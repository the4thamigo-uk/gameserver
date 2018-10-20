package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/the4thamigo_uk/gameserver/pkg/server"
	"os"
)

func main() {
	var cfgPath = pflag.StringP("config", "c", "./config.yaml", "Path or URI to config file in YAML,JSON or TOML format.")
	pflag.Parse()

	cfg, err := server.LoadConfig(*cfgPath)
	if err != nil {
		fmt.Println(err)
		pflag.Usage()
		os.Exit(1)
		return
	}
	s := server.NewServer(cfg)
	err = s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}
