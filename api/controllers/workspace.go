package controllers

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"todo-list/services"
)

type WorkspaceController struct {

}

func (w WorkspaceController) Init(g *echo.Group) {
	g.GET("", getWorkspaces)
	g.GET("/:id", getWorkspace)
	g.POST("", createWorkspace)
	g.DELETE("/:id", deleteWorkspace)
	g.GET("/search", searchWorkspace)
	g.POST("/:id/members", addWorkspaceMember)
}

func getWorkspaces(ctx echo.Context) error {
	uid := userIdFromContext(ctx)
	return services.GetWorkspaces(uid).Send(ctx)
}

func getWorkspace(ctx echo.Context) error {
	uid := userIdFromContext(ctx)
	param := ctx.Param("id")
	wid, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return err
	}
	return services.GetWorkspace(uid, wid).Send(ctx)
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
	param := ctx.Param("id")
	wid, err := strconv.Atoi(param)
	if err != nil {
		return err
	}
	return services.DeleteWorkspace(uid, int64(wid)).Send(ctx)
}

func searchWorkspace(ctx echo.Context) error {
	param := ctx.QueryParam("name")
	return services.SearchWorkspace(userIdFromContext(ctx), param).Send(ctx)
}

func addWorkspaceMember(ctx echo.Context) error {
	uid := userIdFromContext(ctx)
	param := ctx.Param("id")
	wid, err := strconv.Atoi(param)
	if err != nil {
		return err
	}
	return services.AddWorkspaceMember(uid, int64(wid)).Send(ctx)
}