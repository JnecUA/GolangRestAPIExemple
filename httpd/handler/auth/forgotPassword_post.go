package auth

import (
	"fmt"
	"net/http"

	"github.com/JnecUA/GolangRestAPIExemple/platform"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

//ForgotPassword ... request to reset password
func ForgotPassword(smtpConfig map[string]string) gin.HandlerFunc {
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
		//Take a forgot password request
		err = user.ForgotPassword(&conn, smtpConfig)
		if err != nil {
			fmt.Println("Error in user.ForgotPassword()")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": "Reset password mail sended",
		})

	}
}
