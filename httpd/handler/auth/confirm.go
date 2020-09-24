package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

//ConfirmGet ... Confirm account in db
func ConfirmGet(url string) gin.HandlerFunc {
	return func(c *gin.Context) {
		vals := c.Request.URL.Query()
		if vals["hash"] == nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Hash is empty"})
		} else {
			dbpool, err := pgxpool.Connect(context.Background(), url)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"Error": "Unable to connect to database: " + err.Error()})
				return
			}
			defer dbpool.Close()

			//Take Query request to create new User
			var greeting string
			sql := fmt.Sprintf("update users set confirm = true  where randhash = '%v'", vals["hash"])
			_ = dbpool.QueryRow(context.Background(), sql).Scan(&greeting)

			c.JSON(http.StatusOK, gin.H{"Error": "Successfuly"})
		}
	}
}
