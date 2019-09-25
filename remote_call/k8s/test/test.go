package main

import (
	"bytes"
	"exec/pkg/exec"
	"io"
	"k8s.io/client-go/rest"
	"log"
)

func main() {
	//config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	//if err != nil {
	//	log.Fatal(err)
	//}
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
		if err != nil {
			log.Fatalf("error occurs after copying %d bytes: %v", n, err)
		}
	}()

	//err = exec.NewExecutor(config, exec.Stdin(stdin)).ExecCommands("hello-node-78cd77d68f-7scgv", "", []string{"/bin/bash", "-c", `cat | /bin/bash`})
	//if err != nil {
	//	log.Fatal(err)
	//}

	const token = "tpwkoa.ci5xco9mh1tqoblt"
	const url = "https://192.168.99.107:8443"

	config := &rest.Config{
		Host:    url,
		APIPath: "v1",
		ContentConfig: rest.ContentConfig{
			AcceptContentTypes: "application/json",
			ContentType:        "application/json",
		},
		BearerToken: token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}

	err := exec.NewExecutor(config, exec.Stdin(stdin)).ExecScripts("nginx-deployment-547b877857-79k7c", "", []string{
		"/media/storage/github/go_snippet/remote_call/k8s/test/scripts/utils.sh",
		"/media/storage/github/go_snippet/remote_call/k8s/test/scripts/main.sh",
	}, "foo", "bar")
	if err != nil {
		log.Fatal(err)
	}
}
