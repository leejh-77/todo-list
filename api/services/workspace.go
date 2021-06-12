package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"time"
	"todo-list/models"
	"todo-list/orm"
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

	t := ctx.Get("user").(*jwt.Token)
	claims := t.Claims.(jwt.MapClaims)
	uid := claims["uid"].(int64)

	var trError error
	err := orm.InTransaction(func(e orm.Engine) error {
		ws := &models.Workspace{
			Name:        c.Name,
			CreatedTime: time.Now().Unix(),
		}
		id, err := e.Table(models.TableWorkspace).Insert(ws)
		if err != nil {
			trError = err
			return err
		}
		m := &models.WorkspaceMember{
			Type:        models.MemberTypeOwner,
			WorkspaceId: id,
			UserId:      uid,
		}
		_, err = e.Table(models.TableWorkspaceMembers).Insert(m)
		if err != nil {
			trError = err
			return err
		}
		return nil
	})
	if trError != nil {
		return result.ServerError(trError)
	}
	if err != nil {
		return result.ServerError(err)
	}
	return result.Created()
}