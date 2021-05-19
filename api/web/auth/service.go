package auth

import (
	"errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/mail"
	"time"
	"todo-list/model"
	"todo-list/web/result"
)

func signUp(c signUpCommand) *result.ApiResult {
	err := validateSignupRequest(c)
	if err != nil {
		return result.BadRequest(err.Error())
	}
	p, err := encryptPassword(c.Password)
	if err != nil {
		return result.ServerError(err)
	}
	user := new(model.User)
	user.EmailAddress = c.EmailAddress
	user.Password = *p
	user.Username = c.Username
	user.RegisteredTime = time.Now().Unix()

	_, err = createUser(user)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Created()
}

func login(ctx echo.Context, c logInCommand) *result.ApiResult {
	user, err := findUserByEmailAddress(c.EmailAddress)
	if err != nil {
		return result.ServerError(err)
	}
	if user == nil {
		return result.BadRequest("User not found for email - " + c.EmailAddress)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.Password)); err != nil {
		return result.BadRequest("Password not matched")
	}
	token, err := createJwt(c.EmailAddress)
	if err != nil {
		return result.ServerError(err)
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = *token
	cookie.HttpOnly = true
	//cookie.Secure = true
	ctx.SetCookie(cookie)
	return result.Success("")
}

func encryptPassword(p string) (*string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		return nil, errors.New("failed to generate password")
	}
	s := string(hash)
	return &s, nil
}

func validateSignupRequest(c signUpCommand) error {
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