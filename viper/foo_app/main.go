package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	viper.SetEnvPrefix("foo")
	viper.AutomaticEnv()
	os.Setenv("FOO_ID", "123")
	fmt.Println(viper.Get("id"))
}
