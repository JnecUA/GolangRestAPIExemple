package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/JnecUA/GolangRestAPIExemple/platform"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"
)

//UserPut ... Register new user in db
func UserPut(url string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := platform.DefaultUser()

		c.BindJSON(&user)

		validate := validator.New()
		err := validate.Struct(user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": strings.Split(err.Error(), "\n")})
			return
		}

		dbpool, err := pgxpool.Connect(context.Background(), url)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"Error": "Unable to connect to database: " + err.Error()})
		}
		defer dbpool.Close()

		hashedPassword, err := platform.HashPassword(user.Password)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"Error": err})

		}

		var greeting string
		sql := fmt.Sprintf("insert into users (email, username, fullname, passwordhash) values ('%v','%v','%v','%v');", user.Email, user.Username, user.Name, hashedPassword)
		_ = dbpool.QueryRow(context.Background(), sql).Scan(&greeting)
		if greeting == "" {
			c.JSON(http.StatusOK, gin.H{"Status": "Successfuly"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Greeting": greeting})
		}

	}
}
