package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/JnecUA/GolangRestAPIExemple/platform"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

//AuthRequest ... type to proccessing input json data
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//AuthPost ... Auth function
func AuthPost(url string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AuthRequest
		c.BindJSON(&req)

		dbpool, err := pgxpool.Connect(context.Background(), url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		defer dbpool.Close()

		user := platform.DefaultUser()

		sql := fmt.Sprintf("select email, username, fullname, passwordhash, teams_ids, groups_ids from users where username='%s' or email='%s';", req.Username, req.Username)
		err = dbpool.QueryRow(context.Background(), sql).Scan(&user.Email, &user.Username, &user.Name, &user.Password, &user.TeamsIds, &user.GroupIds)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		}

		if user != platform.DefaultUser() {
			match := platform.CheckPasswordHash(req.Password, user.Password)
			if match == true {
				c.JSON(http.StatusOK, gin.H{"user": user})
			} else {
				c.JSON(http.StatusOK, gin.H{"user": "Email or password wrong"})
			}

		} else {
			c.JSON(http.StatusOK, gin.H{"user": "Email or password wrong"})
		}
	}
}
