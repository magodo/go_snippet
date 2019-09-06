package main

import (
	"context"
	"log"
	remote "remote_call/pkg/remote_call"
)

func main() {
	c, err := remote.New(remote.PrivateKeyFile("/home/magodo/.ssh/id_rsa_udb"))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Call(context.Background(), "179.17.0.2", 22,
		[]string{"/media/storage/github/go_snippet/remote_call/ssh/test/scripts/utils.sh",
			"/media/storage/github/go_snippet/remote_call/ssh/test/scripts/main.sh"}, "foo", "bar")
	if err != nil {
		log.Fatal(err)
	}
}
