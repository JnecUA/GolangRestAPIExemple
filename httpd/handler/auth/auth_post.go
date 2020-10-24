package auth

import (
	"net/http"

	"github.com/JnecUA/GolangRestAPIExemple/platform"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

//Auth ... Auth function
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Fill the auth data
		var user platform.User
		c.BindJSON(&user)

		db, _ := c.Get("db")
		conn := db.(pgx.Conn)

		err := user.IsAuthenticated(&conn)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		token, err := user.GetAuthToken()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
			return
		}
	}
}
