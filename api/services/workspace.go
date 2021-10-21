package services

import (
	"errors"
	"time"
	"todo-list/models"
	"todo-list/orm"
	"todo-list/result"
	"todo-list/utils"
)

type CreateWorkspaceCommand struct {
	Name string `json:"name"`
}

type DeleteWorkspaceCommand struct {
	WorkspaceId int64 `json:"workspaceId"`
}

type GetWorkspaceResponse struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	CreatedTime int64 `json:"createdTime"`
	Members []WorkspaceMemberResponse `json:"members"`
	Folders []models.Folder `json:"folders"`
}

type WorkspaceMemberResponse struct {
	UserId int64 `json:"userId"`
	Name string `json:"name"`
	Image *utils.Base64Image `json:"image"`
	Type int `json:"type"`
}

func GetWorkspaces(uid int64) *result.ApiResult {
	var workspaces []models.Workspace
	err := models.WorkspaceQuery(orm.Engine).FindByUserId(&workspaces, uid)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(workspaces)
}

func GetWorkspace(uid int64, wid int64) *result.ApiResult {
	var w models.Workspace
	err := orm.Table(models.TableWorkspace).FindById(&w, wid)
	if err != nil {
		return result.ServerError(err)
	}

	var fs []models.Folder
	err = models.FolderQuery(orm.Engine).FindByWorkspaceId(&fs, wid)
	if err != nil {
		return result.ServerError(err)
	}

	var ms []models.WorkspaceMember
	err = models.WorkspaceMemberQuery(orm.Engine).FindByWorkspaceId(&ms, wid)
	if err != nil {
		return result.ServerError(err)
	}

	res := &GetWorkspaceResponse{}
	res.Id = w.Id
	res.Name = w.Name
	res.CreatedTime = w.CreatedTime
	res.Folders = fs

	members := make([]WorkspaceMemberResponse, 0, len(ms))
	for _, m := range ms {
		var u models.User
		err = orm.Table(models.TableUser).FindById(&u, m.UserId)
		if err != nil {
			return result.ServerError(err)
		}
		mr := WorkspaceMemberResponse{}
		mr.UserId = u.Id
		mr.Name = u.Username
		mr.Type = m.Type
		mr.Image, err = utils.ReadImage(u.Id)
		if err != nil {
			return result.ServerError(err)
		}
		members = append(members, mr)
	}
	res.Members = members
	return result.Success(res)
}

func CreateWorkspace(uid int64, c CreateWorkspaceCommand) *result.ApiResult {
	err := validateCreateWorkspaceCommand(c)
	if err != nil {
		return result.BadRequest(err.Error())
	}

	var ws *models.Workspace
	err = orm.InTransaction(func(e orm.Session) error {
		ws, err = createWorkspace(c, uid, e)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(ws)
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

func SearchWorkspace(uid int64, name string) *result.ApiResult {
	var ws []models.Workspace
	err := models.WorkspaceQuery(orm.Engine).FindByNameLike(&ws, uid, name)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(ws)
}

func validateCreateWorkspaceCommand(c CreateWorkspaceCommand) error {
	if len(c.Name) == 0 {
		return errors.New("name must not be empty")
	}
	return nil
}

func createWorkspace(c CreateWorkspaceCommand, uid int64, e orm.Session) (*models.Workspace, error) {
	ws := &models.Workspace{
		Name:        c.Name,
		CreatedTime: time.Now().Unix(),
	}
	id, err := e.Table(models.TableWorkspace).Insert(ws)
	if err != nil {
		return nil, err
	}
	m := &models.WorkspaceMember{
		Type:        models.MemberTypeOwner,
		WorkspaceId: id,
		UserId:      uid,
	}
	_, err = e.Table(models.TableWorkspaceMember).Insert(m)
	if err != nil {
		return nil, err
	}
	return ws, nil
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
	id, err := tr.Table(models.TableWorkspaceMember).Insert(&m)
	m.Id = id
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