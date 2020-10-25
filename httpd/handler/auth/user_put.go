package auth

import (
	"net/http"
	"strings"

	"github.com/JnecUA/GolangRestAPIExemple/platform"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

//Register ... Register new user in db
func Register(smtpConfig map[string]string) gin.HandlerFunc {
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

		err = user.Register(&conn, smtpConfig)
		if err != nil {
			if len(strings.Split(err.Error(), "\n")) > 1 {
				c.JSON(http.StatusBadRequest, gin.H{"error": strings.Split(err.Error(), "\n")})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": "Account created",
		})
	}
}
