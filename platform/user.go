package platform

import (
	"golang.org/x/crypto/bcrypt"
)

//User ... Struct to work with db
type User struct {
	Email    string   `json:"email" validate:"required,email"`
	Name     string   `json:"name" validate:"required"`
	Username string   `json:"username" validate:"required,gte=5"`
	Password string   `json:"password" validate:"required,gte=8"`
	TeamsIds []string `json:"teamsIds"`
	GroupIds []string `json:"groupIds"`
}

//DefaultUser ... empty user
func DefaultUser() *User {
	return &User{
		Email:    "",
		Name:     "",
		Username: "",
		Password: "",
		TeamsIds: []string{},
		GroupIds: []string{},
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
