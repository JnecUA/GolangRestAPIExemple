package main

import (
	"flag"
	"log"

	"github.com/JnecUA/GolangRestAPIExemple/config"

	"github.com/BurntSushi/toml"
	"github.com/JnecUA/GolangRestAPIExemple/httpd/handler/auth"
	"github.com/gin-gonic/gin"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "./config/apiserver.toml", "path to the toml config file")
}

func main() {
	flag.Parse()

	config := config.DefaultConfig()

	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.POST("/auth", auth.AuthPost(config.DBUrl))
	r.PUT("/auth", auth.UserPut(config.DBUrl))

	r.Run()
}
