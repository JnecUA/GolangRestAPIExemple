package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"

	"github.com/JnecUA/GolangRestAPIExemple/internal/app/apiserver"
	"github.com/JnecUA/GolangRestAPIExemple/internal/app/store"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "path to the toml config file")
}

func main() {
	flag.Parse()
	var GConfig struct {
		ServerConfig map[string]string
		DBConfig     map[string]string
	}
	_, err := toml.DecodeFile(configPath, &GConfig)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.DefaultConfig()
	s.MapConfiguring(GConfig.ServerConfig)
	db := store.DefaultConfig()
	db.MapConfiguring(GConfig.DBConfig)
	App := apiserver.Init(s, db)
	if err := App.Start(); err != nil {
		log.Fatal(err)
	}

}
