package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/JnecUA/GolangRestAPIExemple/config"
	"github.com/jackc/pgx/v4"

	"github.com/BurntSushi/toml"
	"github.com/JnecUA/GolangRestAPIExemple/httpd/handler/auth"
	"github.com/gin-gonic/gin"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "./config/apiserver.toml", "path to the toml config file")
}

func main() {
	//flags parse
	flag.Parse()

	//Empty config
	config := config.DefaultConfig()
	//Filling config
	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := connectDB(config.DBUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Initialize router
	r := gin.Default()

	//Connect DB Middleware
	r.Use(dbMiddleware(*conn))
	//Connect SMTP Middleware
	r.Use(CORSMiddleware())

	//users group
	usersGroup := r.Group("users")
	{
		usersGroup.POST("auth", auth.Auth())
		usersGroup.PUT("register", auth.Register(config.SMTP))
		usersGroup.GET("confirm-account", auth.ConfirmAccount())
		usersGroup.POST("forgot-password", auth.ForgotPassword(config.SMTP))
		usersGroup.POST("reset-password", auth.ResetPassword())
	}

	//Start
	r.Run()
}

func connectDB(url string) (c *pgx.Conn, err error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	_ = conn.Ping(context.Background())
	return conn, err

}

func dbMiddleware(conn pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
