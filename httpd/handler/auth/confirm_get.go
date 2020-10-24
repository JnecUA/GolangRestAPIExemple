package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

//ConfirmAccount ... Confirm account in db
func ConfirmAccount(url string) gin.HandlerFunc {
	return func(c *gin.Context) {
		vals := c.Request.URL.Query()
		if vals["randhash"] == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Hash is empty"})
		} else {
			db, _ := c.Get("db")
			conn := db.(pgx.Conn)

			sql := fmt.Sprintf("update users set confirmed = true where randhash = '%v'", vals["randhash"][0])
			err := conn.QueryRow(context.Background(), sql).Scan()
			if err != pgx.ErrNoRows {
				c.JSON(http.StatusOK, gin.H{"Status": "Successfuly"})
			} else {
				c.JSON(http.StatusOK, gin.H{"Status": "User with this hash not found"})
			}
		}
	}
}
