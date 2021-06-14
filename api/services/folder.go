package services

import (
	"todo-list/models"
	"todo-list/result"
)

type CreateFolderCommand struct {
	WorkspaceId int64
	Name string
}

func GetFolders(uid int64, wid int64) *result.ApiResult {
	ret := checkWorkspaceAuthority(uid, wid)
	if ret != nil {
		return ret
	}
	var fs []models.Folder
	err := models.Folders.FindByWorkspaceId(&fs, wid)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(fs)
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

func DeleteFolder(uid int64, fid int64) *result.ApiResult {
	var f models.Folder
	err := models.Folders.FindById(&f, fid)
	if err != nil {
		return result.ServerError(err)
	}
	if f.Id == int64(0) {
		return result.BadRequest("folder does not exist")
	}
	var w models.Workspace
	err = models.Workspaces.FindById(&w, f.WorkspaceId)
	if err != nil {
		return result.ServerError(err)
	}
	ret := checkWorkspaceAuthority(uid, w.Id)
	if ret != nil {
		return ret
	}
	err = models.Folders.DeleteById(fid)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(nil)
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