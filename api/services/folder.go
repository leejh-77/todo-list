package services

import (
	"todo-list/models"
	"todo-list/result"
)

type CreateFolderCommand struct {
	WorkspaceId int64
	Name string
}

func CreateFolder(uid int64, c CreateFolderCommand) *result.ApiResult {
	ret := checkWorkspaceAuthority(uid, c.WorkspaceId)
	if ret != nil {
		return ret
	}

	f := models.Folder{
		Name:        c.Name,
		WorkspaceId: c.WorkspaceId,
	}
	_, err := models.Folders.Insert(&f)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Created()
}

func checkWorkspaceAuthority(uid int64, wid int64) *result.ApiResult {
	var m models.WorkspaceMember
	err := models.WorkspaceMembers.FindByUserIdAndWorkspaceId(&m, uid, wid)
	if err != nil {
		return result.ServerError(err)
	}

	if m.Id == int64(0) {
		return result.Unauthorized("user is not a member of this workspace")
	}
	return nil
}