package services

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"net/http/httptest"
	"todo-list/models"
)

func createDummyContext() echo.Context {
	return createAuthorizedContext(nil)
}

func createAuthorizedContext(user *models.User) echo.Context {
	req := httptest.NewRequest("GET", "http://localhost", nil)
	if user != nil {
		jwt, err := createJwt(user.Id, user.EmailAddress)
		if err != nil {
			log.Fatal(err)
		}
		req.AddCookie(&http.Cookie{
			Name: "Authorization",
			Value: *jwt,
		})
	}
	rec := httptest.NewRecorder()
	e := echo.New()
	return e.NewContext(req, rec)


}