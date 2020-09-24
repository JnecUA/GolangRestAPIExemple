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
func UserPut(url string, smtp map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Define empty user var
		user := platform.DefaultUser()
		//Fill the variable with data
		c.BindJSON(&user)
		//Validate request data
		validate := validator.New()
		err := validate.Struct(user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": strings.Split(err.Error(), "\n")})
			return
		}
		//Hash password
		hashedPassword, err := platform.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"Error": err})
			return
		}

		//Generate random hash to check email
		randHash := platform.NewSHA1Hash()

		//Connect to DB
		dbpool, err := pgxpool.Connect(context.Background(), url)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"Error": "Unable to connect to database: " + err.Error()})
			return
		}
		defer dbpool.Close()

		//Take Query request to create new User
		var greeting string
		sql := fmt.Sprintf("insert into users (email, username, fullname, passwordhash, randhash, confirmed) values ('%v','%v','%v','%v','%v', false);", user.Email, user.Username, user.Name, hashedPassword, randHash)
		_ = dbpool.QueryRow(context.Background(), sql).Scan(&greeting)

		//Send message to email
		err = platform.SendEmail(smtp, user.Email, randHash)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"Error": err.Error()})
			return
		}

		//Return json
		if greeting == "" {
			c.JSON(http.StatusOK, gin.H{"Status": "Successfuly"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Greeting": greeting})
		}

	}
}
