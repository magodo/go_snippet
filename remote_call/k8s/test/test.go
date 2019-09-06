package main

import (
	"bytes"
	"exec/pkg/exec"
	"io"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	if err != nil {
		log.Fatal(err)
	}
	cmd := bytes.NewBufferString(`
a=1
((a++))
echo "$a"
sleep 5
echo END
`)
	stdin := bytes.NewBuffer([]byte{})

	go func() {
		n, err := io.Copy(stdin, cmd)
		if err != nil{
			log.Fatalf("error occurs after copying %d bytes: %v", n, err)
		}
	}()

	//err = exec.NewExecutor(config, exec.Stdin(stdin)).ExecCommands("hello-node-78cd77d68f-7scgv", "", []string{"/bin/bash", "-c", `cat | /bin/bash`})
	//if err != nil {
	//	log.Fatal(err)
	//}

	err = exec.NewExecutor(config, exec.Stdin(stdin)).ExecScripts("hello-node-78cd77d68f-7scgv", "", []string{
		"/media/storage/github/go_snippet/remote_call/k8s/test/scripts/utils.sh",
		"/media/storage/github/go_snippet/remote_call/k8s/test/scripts/main.sh",
	}, "foo", "bar")
	if err != nil {
		log.Fatal(err)
	}
}


