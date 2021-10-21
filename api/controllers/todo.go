package controllers

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"todo-list/services"
)

type TodoController struct {

}

func (c TodoController) Init(g *echo.Group) {
	g.GET("", getTodos)
	g.POST("", createTodo)
	g.PUT("/:id", updateTodo)
	g.PUT("/positions", updatePositions)
	g.DELETE("/:id", deleteTodo)
}

func getTodos(ctx echo.Context) error {
	param := ctx.QueryParam("folderId")
	fid, err := strconv.Atoi(param)
	if err != nil {
		return err
	}
	return services.GetTodos(userIdFromContext(ctx), int64(fid)).Send(ctx)
}

func createTodo(ctx echo.Context) error {
	uid := userIdFromContext(ctx)
	var c services.CreateTodoCommand
	err := ctx.Bind(&c)
	if err != nil {
		return err
	}
	return services.CreateTodo(uid, c).Send(ctx)
}
func updateTodo(ctx echo.Context) error {
	param := ctx.Param("id")
	tid, err := strconv.Atoi(param)
	if err != nil {
		return err
	}
	uid := userIdFromContext(ctx)
	var c services.UpdateTodoCommand
	err = ctx.Bind(&c)
	if err != nil {
		return err
	}
	c.Id = int64(tid)
	return services.UpdateTodo(uid, c).Send(ctx)
}

func deleteTodo(ctx echo.Context) error {
	param := ctx.Param("id")
	tid, err := strconv.Atoi(param)
	if err != nil {
		return err
	}
	return services.DeleteTodo(userIdFromContext(ctx), int64(tid)).Send(ctx)
}

func updatePositions(ctx echo.Context) error {
	var c services.UpdatePositionsCommand
	err := ctx.Bind(&c)
	if err != nil {
		return err
	}
	return services.UpdatePositions(c).Send(ctx)
}