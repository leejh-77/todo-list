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
	err := models.WorkspaceQuery(orm.Engine).FindByUserId(&workspaces, uid)
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
	err = orm.InTransaction(func(e orm.Session) error {
		return createWorkspace(c, uid, e)
	})
	if err != nil {
		return result.ServerError(err)
	}
	return result.Created()
}

func DeleteWorkspace(uid int64, wid int64) *result.ApiResult {
	var m models.WorkspaceMember
	err := models.WorkspaceMemberQuery(orm.Engine).FindByUserIdAndWorkspaceId(&m, uid, wid)
	if err != nil {
		return result.ServerError(err)
	}
	if m.Id == int64(0) {
		return result.BadRequest("user is not a member of the workspace")
	}
	if m.Type != models.MemberTypeOwner {
		return result.BadRequest("user does not have permission to delete workspace")
	}
	err = orm.InTransaction(func(e orm.Session) error {
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

func createWorkspace(c CreateWorkspaceCommand, uid int64, e orm.Session) error {
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
}

func deleteWorkspace(wid int64, e orm.Session) error {
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

type GetMemberResponseData struct {
	Name string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	IsOwner bool `json:"isOwner"`
}

func GetWorkspaceMembers(uid int64, wid int64) *result.ApiResult {
	var ms []models.WorkspaceMember
	err := models.WorkspaceMemberQuery(orm.Engine).FindByWorkspaceId(&ms, wid)
	if err != nil {
		return result.ServerError(err)
	}

	var ownerId int64
	var isInWorkspace bool
	for _, m := range ms {
		if m.UserId == uid {
			isInWorkspace = true
		}
		if m.Type == models.MemberTypeOwner {
			ownerId = m.UserId
		}
	}

	if !isInWorkspace {
		return result.BadRequest("user is not a member of the workspace")
	}

	var us []models.User
	err = models.UserQuery(orm.Engine).FindByWorkspace(&us, wid)

	res := make([]GetMemberResponseData, 0, 10)
	for _, u := range us {
		data := GetMemberResponseData{
			Name:         u.Username,
			EmailAddress: u.EmailAddress,
			IsOwner: u.Id == ownerId,
		}
		res = append(res, data)
	}
	return result.Success(res)
}

func AddWorkspaceMember(uid int64, wid int64) *result.ApiResult {
	tr, err := orm.BeginTr()
	if err != nil {
		return result.ServerError(err)
	}
	defer tr.Rollback()

	var m models.WorkspaceMember
	err = models.WorkspaceMemberQuery(tr).FindByUserIdAndWorkspaceId(&m, uid, wid)
	if err != nil {
		return result.ServerError(err)
	}
	if m.Id != int64(0) {
		return result.BadRequest("user is already a member of the workspace")
	}
	m = models.WorkspaceMember{
		Type:        models.MemberTypeParticipant,
		WorkspaceId: wid,
		UserId:      uid,
	}
	_, err = tr.Table(models.TableWorkspaceMember).Insert(&m)
	if err != nil {
		return result.ServerError(err)
	}
	err = tr.Commit()
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success("")
}

func DeleteWorkspaceMember(uid int64, wid int64) *result.ApiResult {
	var m models.WorkspaceMember
	err := orm.Table(models.TableWorkspaceMember).Find(&m, "userId = ? AND workspaceId = ?", uid, wid)
	if err != nil {
		return result.ServerError(err)
	}
	if m.Id == int64(0) {
		return result.BadRequest("user is not a member of the workspace")
	}
	err = orm.InTransaction(func(e orm.Session) error {
		return deleteWorkspaceMember(m.Id, wid, e)
	})
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(nil)
}

func deleteWorkspaceMember(mid int64, wid int64, e orm.Session) error {
	err := e.Table(models.TableWorkspaceMember).DeleteById(mid)
	if err != nil {
		return err
	}
	var ms []models.WorkspaceMember
	err = e.Table(models.TableWorkspaceMember).Find(&ms, "workspaceId = ?", wid)
	if err != nil {
		return err
	}
	if len(ms) != 0 {
		return nil
	}
	return e.Table(models.TableWorkspace).DeleteById(wid)
}