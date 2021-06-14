package services

import (
	"errors"
	"time"
	"todo-list/models"
	"todo-list/orm"
	"todo-list/result"
)

type CreateWorkspaceCommand struct {
	Name string
}

func GetWorkspaces(uid int64) *result.ApiResult {
	var workspaces []models.Workspace
	err := models.Workspaces.FindParticipatingWorkspaces(&workspaces, uid)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(workspaces)
}

func CreateWorkspace(uid int64, c CreateWorkspaceCommand) *result.ApiResult {
	err := validateCreateWorkspaceCommand(c)
	if err != nil {
		return result.BadRequest(err.Error())
	}
	err = createWorkspace(c, uid)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Created()
}

func DeleteWorkspace(uid int64, wid int64) *result.ApiResult {
	var m models.WorkspaceMember
	err := models.WorkspaceMembers.FindByUserIdAndWorkspaceId(&m, uid, wid)
	if err != nil {
		return result.ServerError(err)
	}
	if m.Id == int64(0) {
		return result.BadRequest("user is not participating in this workspace")
	}
	if m.Type != models.MemberTypeOwner {
		return result.BadRequest("user does not have permission to delete workspace")
	}
	err = orm.InTransaction(func(e orm.Engine) error {
		return deleteWorkspace(wid, e)
	})
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(nil)
}

func validateCreateWorkspaceCommand(c CreateWorkspaceCommand) error {
	if len(c.Name) == 0 {
		return errors.New("name must not be empty")
	}
	return nil
}

func createWorkspace(c CreateWorkspaceCommand, uid int64) error {
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
		_, err = e.Table(models.TableWorkspaceMember).Insert(m)
		if err != nil {
			return err
		}
		return nil
	})
}

func deleteWorkspace(wid int64, e orm.Engine) error {
	err := e.Table(models.TableWorkspace).DeleteById(wid)
	if err != nil {
		return err
	}
	err = e.Table(models.TableWorkspaceMember).Delete("workspaceId = ?", wid)
	if err != nil {
		return err
	}
	return nil
}
