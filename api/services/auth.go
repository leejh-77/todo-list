package services

import (
	"errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/mail"
	"time"
	"todo-list/base"
	"todo-list/models"
	"todo-list/orm"
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

func SignUp(c SignUpCommand) *result.ApiResult {
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

	_, err = orm.Table(models.TableUser).Insert(user)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Created()
}

func LogIn(ctx echo.Context, c LogInCommand) *result.ApiResult {
	var user models.User
	err := models.UserQuery(orm.Engine).FindByEmailAddress(&user, c.EmailAddress)
	if err != nil {
		return result.ServerError(err)
	}
	if user.Id == int64(0) {
		return result.BadRequest("user not found for email - " + c.EmailAddress)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.Password)); err != nil {
		return result.BadRequest("password not matched")
	}
	token, err := base.CreateJwt(user.Id, c.EmailAddress)
	if err != nil {
		return result.ServerError(err)
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = *token
	cookie.HttpOnly = true
	cookie.MaxAge = 86400
	//cookie.SameSite = http.SameSiteNoneMode
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
