package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"todo-list/constants"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/hello", helloWorld)

	r := e.Group("api")
	r.Use(middleware.JWT(constants.JWTSecret))

	err := e.Start(":9090")
	if err != nil {
		log.Fatal(err)
	}
}

func helloWorld(e echo.Context) error {
	return e.String(http.StatusOK, "hello, world!")
}