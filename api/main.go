package main

import (
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"todo-list/base"
	"todo-list/controllers"
	"todo-list/models"
	"todo-list/orm"
	"todo-list/utils"
)



func initORM() {
	var docker bool
	flag.BoolVar(&docker, "dockerBuild", false, "use docker build")
	flag.Parse()

	if docker {
		base.DBConfig.Host = "host.docker.internal:3306"
	}

	orm.Init(base.DBConfig)
	models.RegisterTables()

	err := utils.CreateImageDir()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initORM()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper:          nil,
		Format:           "[${time_rfc3339_nano}] ${method} ${uri} ${status} ${error}\n",
		CustomTimeFormat: "",
		Output:           nil,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowCredentials: true,
	}))

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
	controllers.UserController{}.Init(r.Group("/users"))

	err := e.Start(":9090")
	if err != nil {
		log.Fatal(err)
	}
}
