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
	r.POST("api/v2/auth", auth.AuthPost(config.DBUrl))
	r.PUT("api/v2/auth", auth.UserPut(config.DBUrl, config.SMTP))
	r.GET("api/v2/confirm-account", auth.ConfirmGet(config.DBUrl))

	r.Run()
}
