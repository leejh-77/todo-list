package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"log"
	"net/http/httptest"
	"time"
	"todo-list/models"
)

func createDummyContext() echo.Context {
	return createAuthorizedContext(nil)
}

func createAuthorizedContext(user *models.User) echo.Context {
	e := echo.New()
	req := httptest.NewRequest("GET", "http://localhost", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	if user != nil {
		c := jwt.MapClaims{}
		c["email"] = user.EmailAddress
		c["uid"] = user.Id
		t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
		ctx.Set("user", t)
	}
	return ctx
}

func testUser() *models.User {
	email := "test.user@gmail.com"

	var user models.User
	err := models.Users.FindByEmailAddress(&user, email)
	if err != nil {
		log.Fatal(err)
	}
	if user.Id != int64(0) {
		return &user
	}
	user = models.User{
		EmailAddress:   email,
		Password:       "paswordl:!@@",
		Username:       "Jonghoon",
		RegisteredTime: time.Now().Unix(),
	}
	id, err := models.Users.Insert(&user)
	if err != nil {
		log.Fatal(err)
	}
	user.Id = id
	return &user
}