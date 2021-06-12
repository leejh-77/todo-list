package controllers

import (
	"github.com/labstack/echo/v4"
	"todo-list/services"
)

type TodoController struct {

}

func (c TodoController) Init(g *echo.Group) {
	g.POST("", createTodo, withCommand(services.CreateTodoCommand{}))
}

func createTodo(cxt echo.Context) error {
	return send(cxt, services.CreateTodo(cxt))
}