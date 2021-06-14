package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http/httptest"
	"todo-list/models"
)

func clearTables() {
	_ = models.Folders.DeleteAll()
	_ = models.Todos.DeleteAll()
	_ = models.Users.DeleteAll()
	_ = models.Workspaces.DeleteAll()
	_ = models.WorkspaceMembers.DeleteAll()
}

func createDummyContext() echo.Context {
	return createContext(nil)
}

func createAuthorizedContext() echo.Context {
	return createContext(models.TestUser())
}

func createContext(user *models.User) echo.Context {
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