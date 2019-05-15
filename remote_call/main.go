package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {

	host := "localhost"
	port := "32804"
	user := "root"
	script := "scripts/foo.sh"
	args := []string{"foo", "1 2"}

	key, err := ioutil.ReadFile("/home/magodo/.ssh/id_rsa_foo")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse public key: %v", err)
	}

	// ssh client config
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		// allow any host key to be used (non-prod)
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),

		// optional host key algo list
		HostKeyAlgorithms: []string{
			ssh.KeyAlgoRSA,
			ssh.KeyAlgoDSA,
			ssh.KeyAlgoECDSA256,
			ssh.KeyAlgoECDSA384,
			ssh.KeyAlgoECDSA521,
			ssh.KeyAlgoED25519,
		},
		// optional tcp connect timeout
		Timeout: 5 * time.Second,
	}

	if host != "localhost" {
		// get host public key
		hostKey := getHostKey(host)
		// verify host public key
		config.HostKeyCallback = ssh.FixedHostKey(hostKey)
	}

	// connect
	client, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// start session
	sess, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	// setup standard out and error
	// uses writer interface
	cmdIn, err := sess.StdinPipe()
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	// run single command
	err = sess.Shell()
	if err != nil {
		log.Fatal(err)
	}

	scriptFile, err := os.Open(script)
	if err != nil {
		log.Fatal(err)
	}
	defer scriptFile.Close()

	go func() {
		defer func() {
			lastCmd := "main"
			for _, arg := range args {
				lastCmd += " \"" + arg + "\""
			}
			cmdIn.Write([]byte(lastCmd))
			cmdIn.Close()
		}()

		firstCmd := "export ACTION_MODE=on\n"
		cmdIn.Write([]byte(firstCmd))
		_, err = io.Copy(cmdIn, scriptFile)
		if err != nil {
			log.Fatal(err)
		}

	}()

	err = sess.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func getHostKey(host string) ssh.PublicKey {
	// parse OpenSSH known_hosts file
	// ssh or use ssh-keyscan to get initial key
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				log.Fatalf("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}

	if hostKey == nil {
		log.Fatalf("no hostkey found for %s", host)
	}

	return hostKey
}
