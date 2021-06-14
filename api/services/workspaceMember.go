package services

import (
	"todo-list/models"
	"todo-list/orm"
	"todo-list/result"
)

type GetUserResponseData struct {
	Name string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	IsOwner bool `json:"isOwner"`
}

func GetWorkspaceMembers(uid int64, wid int64) *result.ApiResult {
	var ms []models.WorkspaceMember
	err := models.WorkspaceMembers.FindByWorkspace(&ms, wid)
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
	err = models.Users.FindByWorkspace(&us, wid)

	res := make([]GetUserResponseData, 0, 10)
	for _, u := range us {
		data := GetUserResponseData{
			Name:         u.Username,
			EmailAddress: u.EmailAddress,
			IsOwner: u.Id == ownerId,
		}
		res = append(res, data)
	}
	return result.Success(res)
}

func DeleteWorkspaceMember(uid int64, wid int64) *result.ApiResult {
	var m models.WorkspaceMember
	err := models.WorkspaceMembers.FindByUserIdAndWorkspaceId(&m, uid, wid)
	if err != nil {
		return result.ServerError(err)
	}
	if m.Id == int64(0) {
		return result.BadRequest("user is not a member of the workspace")
	}
	err = deleteWorkspaceMember(m.Id, wid)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(nil)
}

func deleteWorkspaceMember(mid int64, wid int64) error {
	return orm.InTransaction(func(e orm.Engine) error {
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
		return deleteWorkspace(wid, e)
	})
}