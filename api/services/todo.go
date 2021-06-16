package services

import (
	"errors"
	"todo-list/models"
	"todo-list/orm"
	"todo-list/result"
)

type CreateTodoCommand struct {
	FolderId int64 `json:"folderIdd"`
	Subject string `json:"subject"`
	Body string `json:"body"`
	Status int `json:"status"`
	CompletedTime int64 `json:"completedTime"`
}

func GetTodo(uid int64, fid int64) *result.ApiResult {
	var m models.WorkspaceMember
	err := models.WorkspaceMemberQuery(orm.Engine).FindByUserIdAndFolderId(&m, uid, fid)
	if err != nil {
		return result.ServerError(err)
	}
	if m.Id == int64(0) {
		return result.BadRequest("user is not a member of the workspace of the folder")
	}
	var fs []models.Todo
	err = models.TodoQuery(orm.Engine).FindByFolderId(&fs, fid)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(fs)
}

func CreateTodo(uid int64, c CreateTodoCommand) *result.ApiResult {
	err := validateCreateCommand(c)
	if err != nil {
		return result.BadRequest(err.Error())
	}
	var m models.WorkspaceMember
	err = models.WorkspaceMemberQuery(orm.Engine).FindByUserIdAndFolderId(&m, uid, c.FolderId)
	if err != nil {
		return result.ServerError(err)
	}
	if m.Id == int64(0) {
		return result.BadRequest("user is not a member of the workspace of the folder")
	}
	todo := &models.Todo{
		FolderId:      c.FolderId,
		UserId: uid,
		Subject:       c.Subject,
		Body:          c.Body,
		Status:        c.Status,
		CompletedTime: c.CompletedTime,
		Position:      0,
	}
	_, err = orm.Table(models.TableTodo).Insert(todo)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Created()
}

func validateCreateCommand(c CreateTodoCommand) error {
	if c.FolderId == int64(0) {
		return errors.New("folder id must not be empty")
	}
	if c.Status != models.TodoStatusCompleted && c.CompletedTime != int64(0) {
		return errors.New("uncompleted todo cannot have property `completedTime`")
	}
	return nil
}