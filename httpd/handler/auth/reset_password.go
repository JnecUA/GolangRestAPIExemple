package auth

import (
	"fmt"
	"net/http"

	"github.com/JnecUA/GolangRestAPIExemple/platform"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

//ResetPassword set new password
func ResetPassword() gin.HandlerFunc {
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
		err = user.ResetPassword(&conn)
		if err != nil {
			fmt.Println("Error in user.ResetPassword()")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": "Password successfully reset",
		})
	}

}
