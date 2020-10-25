package platform

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	tokenSecret = []byte(os.Getenv("TOKEN_SECRET"))
)

//User ... Struct to work with db
type User struct {
	Email           string   `json:"email" validate:"required,email"`
	Name            string   `json:"name" validate:"required"`
	Username        string   `json:"username" validate:"required,gte=5"`
	PasswordHash    string   `json:"-"`
	Password        string   `json:"password" validate:"required,gte=8"`
	PasswordConfirm string   `json:"password_confirm" validate:"required,gte=8"`
	TeamsIds        []string `json:"teamsIds"`
	GroupIds        []string `json:"groupIds"`
	Confirmed       bool     `json:"confirmed"`
	RandHash        string   `json:"randhash"`
}

//Register ... Registrate new user in DB
func (u *User) Register(conn *pgx.Conn, smtpConfig map[string]string) error {
	//We prepare sql request
	sql := fmt.Sprintf("select email from users where username='%s' or email='%s';", u.Email, u.Email)

	row := conn.QueryRow(context.Background(), sql)

	userLookup := User{}
	err := row.Scan(&userLookup)
	if err != pgx.ErrNoRows {
		return fmt.Errorf("A user with than email or username already exist")
	}

	//Validate request data
	validate := validator.New()
	err = validate.Struct(u)
	if err != nil {
		return err
	}

	//Hash password
	hashedPassword, err := HashPassword(u.Password)
	if err != nil {
		return err
	}

	//Generate random hash to check email
	randHash := NewSHA1Hash()

	//Take Query request to create new User
	sql = fmt.Sprintf("insert into users (email, username, fullname, passwordhash, randhash) values ('%v','%v','%v','%v','%v');", strings.ToLower(u.Email), u.Username, u.Name, hashedPassword, randHash)
	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		return err
	}
	//Send message to email

	err = SendConfirmEmail(smtpConfig, u.Email, randHash)

	return err
}

//IsAuthenticated ... User authorization check
func (u *User) IsAuthenticated(conn *pgx.Conn) error {
	//Query request and fill user data
	sql := fmt.Sprintf("select email, username, fullname, passwordhash, teams_ids, groups_ids from users where username='%s' or email='%s';", u.Email, u.Email)
	err := conn.QueryRow(context.Background(), sql).Scan(&u.Email, &u.Username, &u.Name, &u.PasswordHash, &u.TeamsIds, &u.GroupIds)
	if err == pgx.ErrNoRows {
		return fmt.Errorf("User with email or username not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password))
	if err != nil {
		return fmt.Errorf("Password wrong")
	}
	return nil
}

//ForgotPassword request to reset password
func (u *User) ForgotPassword(conn *pgx.Conn, smtpConfig map[string]string) error {
	//Generate random hash to check email
	randHash := NewSHA1Hash()

	sql := fmt.Sprintf("update users set randhash = '%v' where email='%s';", randHash, u.Email)
	res, err := conn.Exec(context.Background(), sql)
	if err != nil {
		return fmt.Errorf("Error on sending query request")
	}
	if string(res) == "UPDATE 0" {
		return fmt.Errorf("Email does not exist")
	}
	err = SendResetEmail(smtpConfig, u.Email, randHash)
	if err != nil {
		return nil
	}
	return nil
}

//GetAuthToken take a jwt token
func (u *User) GetAuthToken() (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	authToken, err := token.SignedString(tokenSecret)
	return authToken, err
}

//ResetPassword set new password
func (u *User) ResetPassword(conn *pgx.Conn) error {
	if u.Password != u.PasswordConfirm {
		return fmt.Errorf("Fields password and password_confirm does not match")
	}
	//Generate new password
	newPass, err := HashPassword(u.Password)
	if err != nil {
		return err
	}
	sql := fmt.Sprintf("update users set passwordhash = '%v' where randhash = '%v'", newPass, u.RandHash)
	res, err := conn.Exec(context.Background(), sql)
	if err != nil {
		return fmt.Errorf("Error on sending query request")
	}
	if string(res) == "UPDATE 0" {
		return fmt.Errorf("User with this hash not found")
	}
	return nil
}

//HashPassword ... Hashing password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

//SendConfirmEmail ... Send message to email address
func SendConfirmEmail(smtpConf map[string]string, receiver string, hash string) error {
	auth := sasl.NewPlainClient("", smtpConf["user"], smtpConf["pass"])
	// Connect to the server, authenticate, set the sender and receiver,
	// and send the email all in one step.
	to := []string{receiver}
	msg := strings.NewReader("To: " + receiver + "\r\n" +
		"Subject: Confirm your email!\r\n" +
		"\r\n" +
		"Click on link to confirm your account localhost:8080/users/confirm-account?randhash=" + hash)
	err := smtp.SendMail(smtpConf["host"]+":"+smtpConf["port"], auth, smtpConf["sender"], to, msg)
	if err != nil {
		return err
	}
	return nil
}

//SendResetEmail ... Send eemail to reset password
func SendResetEmail(smtpConf map[string]string, receiver string, hash string) error {
	auth := sasl.NewPlainClient("", smtpConf["user"], smtpConf["pass"])
	// Connect to the server, authenticate, set the sender and receiver,
	// and send the email all in one step.
	to := []string{receiver}
	msg := strings.NewReader("To: " + receiver + "\r\n" +
		"Subject: Password Reset\r\n" +
		"\r\n" +
		"Click on link to reset your password *Your client domain*/users/reset-password?randhash=" + hash)
	err := smtp.SendMail(smtpConf["host"]+":"+smtpConf["port"], auth, smtpConf["sender"], to, msg)
	if err != nil {
		return err
	}
	return nil
}
