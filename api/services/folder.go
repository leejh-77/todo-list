package services

import (
	"todo-list/models"
	"todo-list/orm"
	"todo-list/result"
)

type CreateFolderCommand struct {
	WorkspaceId int64 `json:"workspaceId"`
	Name string `json:"name"`
}

type DeleteFolderCommand struct {
	FolderId int64
}

func GetFolders(uid int64, wid int64) *result.ApiResult {
	ret := checkWorkspaceAuthority(uid, wid)
	if ret != nil {
		return ret
	}
	var fs []models.Folder
	err := orm.Table(models.TableFolder).Find(&fs, "workspaceId = ?", wid)
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
	id, err := orm.Table(models.TableFolder).Insert(&f)
	f.Id = id
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(f)
}

func DeleteFolder(uid int64, fid int64) *result.ApiResult {
	var m models.WorkspaceMember
	err := models.WorkspaceMemberQuery(orm.Engine).FindByUserIdAndFolderId(&m, uid, fid)
	if err != nil {
		return result.ServerError(err)
	}
	if m.Id == int64(0) {
		return result.BadRequest("user is not a member of the workspace")
	}
	err = orm.Table(models.TableFolder).DeleteById(fid)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(nil)
}

func checkWorkspaceAuthority(uid int64, wid int64) *result.ApiResult {
	var m models.WorkspaceMember
	err := models.WorkspaceMemberQuery(orm.Engine).FindByUserIdAndWorkspaceId(&m, uid, wid)
	if err != nil {
		return result.ServerError(err)
	}
	if m.Id == int64(0) {
		return result.BadRequest("user is not a member of the workspace")
	}
	return nil
}