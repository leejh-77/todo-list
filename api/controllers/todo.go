package controllers

import (
	"github.com/labstack/echo/v4"
	"todo-list/services"
)

type TodoController struct {

}

func (c TodoController) Init(g *echo.Group) {
	g.POST("", createTodo)
}

func createTodo(ctx echo.Context) error {
	uid := userIdFromContext(ctx)
	c := services.CreateTodoCommand{}
	err := ctx.Bind(&c)
	if err != nil {
		return err
	}
	return services.CreateTodo(uid, c).Send(ctx)
}