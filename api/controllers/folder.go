package controllers

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"todo-list/services"
)

type FolderController struct {

}

func (f FolderController) Init(g *echo.Group) {
	g.GET("", getFolders)
	g.POST("", addFolder)
	g.DELETE("/:id", deleteFolder)
}

func getFolders(ctx echo.Context) error {
	param := ctx.QueryParam("workspaceId")
	wid, err := strconv.Atoi(param)
	if err != nil {
		return err
	}
	return services.GetFolders(userIdFromContext(ctx), int64(wid)).Send(ctx)
}

func addFolder(ctx echo.Context) error {
	var c services.CreateFolderCommand
	err := ctx.Bind(&c)
	if err != nil {
		return err
	}
	return services.CreateFolder(userIdFromContext(ctx), c).Send(ctx)
}

func deleteFolder(ctx echo.Context) error {
	param := ctx.Param("id")
	fid, err := strconv.Atoi(param)
	if err != nil {
		return err
	}
	return services.DeleteFolder(userIdFromContext(ctx), int64(fid)).Send(ctx)
}