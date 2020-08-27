package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"

	"github.com/JnecUA/GolangRestAPIExemple/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "path to the toml config file")
}

func main() {
	flag.Parse()
	config := apiserver.DefaultConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	App := apiserver.Init(config)
	if err := App.Start(); err != nil {
		log.Fatal(err)
	}

}
