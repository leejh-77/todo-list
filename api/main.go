package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"todo-list/base"
	"todo-list/controllers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/hello", helloWorld)

	controllers.AuthController{}.Init(e.Group("/"))

	r := e.Group("api")
	r.Use(middleware.JWT(base.JWTSecret))
	controllers.WorkspaceController{}.Init(r.Group("/workspaces"))
	controllers.FolderController{}.Init(r.Group("/folders"))
	controllers.TodoController{}.Init(r.Group("/todos"))

	err := e.Start(":9090")
	if err != nil {
		log.Fatal(err)
	}
}

func helloWorld(e echo.Context) error {
	return e.String(http.StatusOK, "hello, world!")
}