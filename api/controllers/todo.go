package controllers

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"todo-list/services"
)

type TodoController struct {

}

func (c TodoController) Init(g *echo.Group) {
	g.POST("", createTodo)
	g.GET("", getTodos)
	g.PUT("/:id/position", moveTodo)
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

func getTodos(ctx echo.Context) error {
	param := ctx.QueryParam("folderId")
	fid, err := strconv.Atoi(param)
	if err != nil {
		return err
	}
	return services.GetTodos(userIdFromContext(ctx), int64(fid)).Send(ctx)
}

func moveTodo(ctx echo.Context) error {
	param := ctx.Param("id")
	tid, err := strconv.Atoi(param)
	if err != nil {
		return err
	}
	var c services.MoveTodoCommand
	err = ctx.Bind(&c)
	if err != nil {
		return err
	}
	return services.MoveTodo(userIdFromContext(ctx), int64(tid), c).Send(ctx)
}