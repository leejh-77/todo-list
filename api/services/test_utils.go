package services

import (
	"github.com/labstack/echo/v4"
	"net/http/httptest"
)

func dummyContext() echo.Context {
	req := httptest.NewRequest("GET", "http://localhost", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	return e.NewContext(req, rec)
}