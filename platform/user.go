package platform

import (
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"golang.org/x/crypto/bcrypt"
)

//User ... Struct to work with db
type User struct {
	Email     string   `json:"email" validate:"required,email"`
	Name      string   `json:"name" validate:"required"`
	Username  string   `json:"username" validate:"required,gte=5"`
	Password  string   `json:"password" validate:"required,gte=8"`
	TeamsIds  []string `json:"teamsIds"`
	GroupIds  []string `json:"groupIds"`
	Confirmed bool     `json:"confirmed"`
	RandHash  string   `json:"randhash"`
}

//DefaultUser ... empty user
func DefaultUser() *User {
	return &User{
		Email:     "",
		Name:      "",
		Username:  "",
		Password:  "",
		TeamsIds:  []string{},
		GroupIds:  []string{},
		Confirmed: false,
		RandHash:  "",
	}
}

//HashPassword ... Hashing password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash ... Compare passwords from db and request
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//SendEmail ... Send message to email address
func SendEmail(smtpConf map[string]string, receiver string, hash string) error {
	auth := sasl.NewPlainClient("", smtpConf["user"], smtpConf["pass"])
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{receiver}
	msg := strings.NewReader("To: " + receiver + "\r\n" +
		"Subject: Confirm your account!\r\n" +
		"\r\n" +
		`Click on link to confirm your account <a href="localhost:8080/api/v2/confirm-account?randhash=` + hash + `">Confirm</a>.\r\n`)
	err := smtp.SendMail(smtpConf["host"]+":"+smtpConf["port"], auth, smtpConf["sender"], to, msg)
	if err != nil {
		return err
	}
	return nil
}
