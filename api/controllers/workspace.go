package controllers

import (
	"github.com/labstack/echo/v4"
	"todo-list/services"
)

type WorkspaceController struct {

}

func (w WorkspaceController) Init(g *echo.Group) {
	g.GET("", getWorkspaces)
	g.POST("", createWorkspace)
	g.DELETE("/:id", deleteWorkspace)
}

func getWorkspaces(ctx echo.Context) error {
	uid := userIdFromContext(ctx)
	return services.GetWorkspaces(uid).Send(ctx)
}

func createWorkspace(ctx echo.Context) error {
	uid := userIdFromContext(ctx)
	var c services.CreateWorkspaceCommand
	err := ctx.Bind(&c)
	if err != nil {
		return err
	}
	return services.CreateWorkspace(uid, c).Send(ctx)
}

func deleteWorkspace(ctx echo.Context) error {
	uid := userIdFromContext(ctx)
	var c services.DeleteWorkspaceCommand
	err := ctx.Bind(&c)
	if err != nil {
		return err
	}
	return services.DeleteWorkspace(uid, c).Send(ctx)
}