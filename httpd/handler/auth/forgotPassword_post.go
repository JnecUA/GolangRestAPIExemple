package auth

import (
	"net/http"

	"github.com/JnecUA/GolangRestAPIExemple/platform"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

//ForgotPassword ... request to reset password
func ForgotPassword(url string, smtp map[string]string) gin.HandlerFunc {
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
		//Take a forgot password request
		user.ForgotPassword(&conn, smtpConfig)

	}
}
