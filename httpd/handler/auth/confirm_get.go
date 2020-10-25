package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

//ConfirmAccount ... Confirm account in db
func ConfirmAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		vals := c.Request.URL.Query()
		if vals["randhash"] == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Hash is empty"})
		} else {
			db, _ := c.Get("db")
			conn := db.(pgx.Conn)
			sql := fmt.Sprintf("update users set confirmed = true where randhash = '%v'", vals["randhash"][0])
			res, err := conn.Exec(context.Background(), sql)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": "Error on sending query request"})
			}

			if string(res) == "UPDATE 0" {
				c.JSON(http.StatusOK, gin.H{"error": "User with this hash not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": "Confirmed"})
		}
	}
}
