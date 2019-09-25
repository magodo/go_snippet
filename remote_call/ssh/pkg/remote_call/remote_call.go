package remote_call

import (
	"context"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

type Option func(*Options)

type RemoteCall interface {
	Call(ctx context.Context, host string, port int, scripts []string, args ...string) error // start a new session to invoke a remote call, then close it
}

type remoteCall struct {
	*ssh.ClientConfig
	*Options
}

// New create a new RemoteCall object.
func New(opts ...Option) (RemoteCall, error) {
	c := &remoteCall{}

	// customize options
	options := newDefaultOptions()
	for _, opt := range opts {
		opt(options)
	}
	c.Options = options

	key, err := ioutil.ReadFile(options.privateKeyFile)
	if err != nil {
		errors.Errorf("unable to read private key: %v", err)
		return c, err
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		errors.Errorf("unable to parse public key: %v", err)
		return c, err
	}

	// ssh client config
	config := &ssh.ClientConfig{
		User: options.user,
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
		Timeout: options.timeout,
	}
	c.ClientConfig = config
	return c, nil
}

// Take care of the order in `scripts` argument, it is executed in the same order as passed.
func (c *remoteCall) Call(ctx context.Context, host string, port int, scripts []string, args ...string) error {
	client, err := ssh.Dial("tcp", net.JoinHostPort(host, strconv.Itoa(port)), c.ClientConfig)
	if err != nil {
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		err = errors.Wrap(err, "failed to new session")
		return err
	}
	defer session.Close()

	cmdIn, err := session.StdinPipe()
	session.Stdout = c.Options.stdout
	session.Stderr = c.Options.stderr

	err = session.Shell()
	if err != nil {
		err = errors.Wrap(err, "failed to login shell")
		return err
	}

	mergedScript := io.MultiReader()
	for _, script := range scripts {
		f, err := os.Open(script)
		if err != nil {
			err = errors.Wrapf(err, "failed to open script: %s", script)
			return err
		}
		defer f.Close()
		mergedScript = io.MultiReader(mergedScript, f)
	}

	var errCmd error
	go func() {
		defer func() {
			if errCmd == nil {
				lastCmd := "main"
				for _, arg := range args {
					lastCmd += " '" + arg + "'"
				}
				_, errCmd = cmdIn.Write([]byte(lastCmd))
				if errCmd != nil {
					return
				}
				errCmd = cmdIn.Close()
			}
		}()
		firstCmd := "export REMOTE_MODE=on"
		_, errCmd = cmdIn.Write([]byte(firstCmd))
		if errCmd != nil {
			return
		}
		_, errCmd = io.Copy(cmdIn, mergedScript)
	}()

	err = session.Wait()
	if err != nil {
		return err
	}

	if errCmd != nil {
		err = errors.Wrap(err, "an error occured when feeding command to remote shell")
		return err
	}

	return nil
}
