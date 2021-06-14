package services

import (
	"errors"
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
	err := validateCreateWorkspaceCommand(c)
	if err != nil {
		return result.BadRequest(err.Error())
	}

	t := ctx.Get("user").(*jwt.Token)
	claims := t.Claims.(jwt.MapClaims)
	uid := claims["uid"].(int64)

	err = createWorkspaceAndAddMember(c, uid)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Created()
}

func validateCreateWorkspaceCommand(c CreateWorkspaceCommand) error {
	if len(c.Name) == 0 {
		return errors.New("name must not be empty")
	}
	return nil
}

func createWorkspaceAndAddMember(c CreateWorkspaceCommand, uid int64) error {
	return orm.InTransaction(func(e orm.Engine) error {
		ws := &models.Workspace{
			Name:        c.Name,
			CreatedTime: time.Now().Unix(),
		}
		id, err := e.Table(models.TableWorkspace).Insert(ws)
		if err != nil {
			return err
		}
		m := &models.WorkspaceMember{
			Type:        models.MemberTypeOwner,
			WorkspaceId: id,
			UserId:      uid,
		}
		_, err = e.Table(models.TableWorkspaceMembers).Insert(m)
		if err != nil {
			return err
		}
		return nil
	})
}
