package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
)

func main() {
	var server = pflag.StringP("server", "s", "http://127.0.0.1:8080", "Root uri of game server instance.")
	pflag.Parse()

	err := run(*server)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
