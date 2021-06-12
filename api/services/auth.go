package services

import (
	"errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/mail"
	"time"
	"todo-list/models"
	"todo-list/result"
)

type SignUpCommand struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
	Username     string `json:"username"`
}

type LogInCommand struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

func SignUp(ctx echo.Context) *result.ApiResult {
	c := ctx.Get("command").(SignUpCommand)

	err := validateSignupRequest(c)
	if err != nil {
		return result.BadRequest(err.Error())
	}
	p, err := encryptPassword(c.Password)
	if err != nil {
		return result.ServerError(err)
	}
	user := new(models.User)
	user.EmailAddress = c.EmailAddress
	user.Password = *p
	user.Username = c.Username
	user.RegisteredTime = time.Now().Unix()

	_, err = models.UserTable.Insert(user)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Created()
}

func LogIn(ctx echo.Context) *result.ApiResult {
	c := ctx.Get("command").(LogInCommand)

	var user models.User
	err := models.UserTable.FindByEmailAddress(&user, c.EmailAddress)
	if err != nil {
		return result.ServerError(err)
	}
	if user.Id == int64(0) {
		return result.BadRequest("User not found for email - " + c.EmailAddress)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.Password)); err != nil {
		return result.BadRequest("Password not matched")
	}
	token, err := createJwt(user.Id, c.EmailAddress)
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
