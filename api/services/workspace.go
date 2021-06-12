package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"time"
	"todo-list/models"
	"todo-list/result"
)

type CreateWorkspaceCommand struct {
	Name string
}

func CreateWorkspace(ctx echo.Context) *result.ApiResult {
	c := ctx.Get("command").(CreateWorkspaceCommand)
	if len(c.Name) == 0 {
		return result.BadRequest("name must not be empty")
	}
	ws := &models.Workspace{
		Name:        c.Name,
		CreatedTime: time.Now().Unix(),
	}
	id, err := models.Workspaces.Insert(ws)
	if err != nil {
		return result.ServerError(err)
	}
	t := ctx.Get("user").(*jwt.Token)
	claims := t.Claims.(jwt.MapClaims)
	uid := claims["uid"].(int64)

	m := &models.WorkspaceMember{
		Type:        models.MemberTypeOwner,
		WorkspaceId: id,
		UserId:      uid,
	}
	_, err = models.WorkspaceMembers.Insert(m)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Created()
}