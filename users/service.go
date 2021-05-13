package users

import (
	"errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"time"
	"todo-list/auth"
	"todo-list/model"
)

type SignUpCommand struct {
	EmailAddress string
	Password string
	Username string
}

func SignUp(cxt echo.Context, c SignUpCommand) (*string, error) {
	err := validateSignupRequest(c)
	if err != nil {
		return nil, err
	}

	encrypted, err := encryptPassword(c.Password)
	if err != nil {
		return nil, err
	}

	user := new(model.User)
	user.EmailAddress = c.EmailAddress
	user.Password = *encrypted
	user.Username = c.Username
	user.RegisteredTime = time.Now().Unix()

	err = createUser(user)
	if err != nil {
		return nil, err
	}
	token, err := auth.CreateJwt(c.EmailAddress)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func encryptPassword(p string) (*string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		return nil, err
	}
	s := string(hash)
	return &s, nil
}

func validateSignupRequest(c SignUpCommand) error {
	email := c.EmailAddress
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid emailAddress")
	}
	password := c.Password
	if len(password) <= 4 {
		return errors.New("invalid password")
	}
	username := c.Username
	if len(username) <= 4 {
		return errors.New("invalid username")
	}
	return nil
}