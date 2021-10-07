package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"todo-list/base"
	"todo-list/controllers"
	"todo-list/models"
	"todo-list/orm"
)

func initORM() {
	orm.Init(base.DBConfig)
	models.RegisterTables()
}

func main() {
	initORM()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowCredentials: true,
	}))
	e.GET("/hello", helloWorld)

	controllers.AuthController{}.Init(e.Group(""))

	r := e.Group("api")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:              []byte(base.JWTSecret),
		SigningMethod:           "HS256",
		TokenLookup:             "cookie:token",
	}))

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