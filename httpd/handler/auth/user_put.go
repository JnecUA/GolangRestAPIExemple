package auth

import (
	"fmt"
	"net/http"

	"github.com/JnecUA/GolangRestAPIExemple/platform"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

//Register ... Register new user in db
func Register(url string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Filling data
		user := platform.User{}
		err := c.ShouldBindJSON(&user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//Get db connect var
		db, _ := c.Get("db")
		conn := db.(pgx.Conn)
		//Get smtp config var
		smtp, _ := c.Get("smtp")
		smtpConfig := smtp.(map[string]string)
		//User Register
		err = user.Register(&conn, smtpConfig)
		if err != nil {
			fmt.Println("Error in user.Register()")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": "Account created",
		})
	}
}
