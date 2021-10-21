package services

import (
	"errors"
	"todo-list/models"
	"todo-list/optional"
	"todo-list/orm"
	"todo-list/result"
)

type CreateTodoCommand struct {
	FolderId      int64   `json:"folderId"`
	Subject       string  `json:"subject"`
	Body          string  `json:"body"`
	Status        int     `json:"status"`
	CompletedTime int64   `json:"completedTime"`
	Position      float32 `json:"position"`
}

type UpdateTodoCommand struct {
	Id            int64           `json:"id"`
	Subject       optional.String `json:"subject"`
	Body          optional.String `json:"body"`
	Status        optional.Int    `json:"status"`
	CompletedTime optional.Int64  `json:"completedTime"`
}

type MoveTodoCommand struct {
	FolderId int64
	Status   optional.Int `json:"status"`
	Position float32      `json:"position"`
}

func (c *UpdateTodoCommand) hasChange() bool {
	return c.Subject.Set || c.Body.Set || c.Status.Set || c.CompletedTime.Set
}

func GetTodos(uid int64, fid int64) *result.ApiResult {
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
	ret := checkFolderAuthority(uid, c.FolderId)
	if ret != nil {
		return ret
	}
	todo := &models.Todo{
		FolderId:      c.FolderId,
		UserId:        uid,
		Subject:       c.Subject,
		Body:          c.Body,
		Status:        c.Status,
		CompletedTime: c.CompletedTime,
		Position:      c.Position,
	}
	id, err := orm.Table(models.TableTodo).Insert(todo)
	todo.Id = id
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(todo)
}

func UpdateTodo(uid int64, c UpdateTodoCommand) *result.ApiResult {
	if c.Id == int64(0) {
		return result.BadRequest("id must not be empty")
	}
	t, ret := findTodo(c.Id)
	if ret != nil {
		return ret
	}
	ret = checkFolderAuthority(uid, t.FolderId)
	if ret != nil {
		return ret
	}
	if !c.hasChange() {
		return result.Success("")
	}
	if c.Subject.Set {
		t.Subject = c.Subject.Value
	}
	if c.Body.Set {
		t.Body = c.Body.Value
	}
	if c.Status.Set {
		t.Status = c.Status.Value
	}
	if c.CompletedTime.Set {
		t.CompletedTime = c.CompletedTime.Value
	}
	if t.Status != models.TodoStatusCompleted && t.CompletedTime != int64(0) {
		return result.BadRequest("uncompleted todo cannot have property `completedTime`")
	}
	err := orm.Table(models.TableTodo).Update(t)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(t)
}

func DeleteTodo(uid int64, tid int64) *result.ApiResult {
	if tid == int64(0) {
		return result.BadRequest("id must not be empty")
	}
	t, ret := findTodo(tid)
	if ret != nil {
		return ret
	}
	ret = checkFolderAuthority(uid, t.FolderId)
	if ret != nil {
		return ret
	}
	err := orm.Table(models.TableTodo).DeleteById(t.Id)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success("")
}

func MoveTodo(uid int64, tid int64, c MoveTodoCommand) *result.ApiResult {
	if tid == int64(0) {
		return result.BadRequest("todo id must not be empty")
	}
	t, ret := findTodo(tid)
	if ret != nil {
		return ret
	}
	ret = checkFolderAuthority(uid, t.FolderId)
	if ret != nil {
		return ret
	}
	var todo models.Todo
	err := orm.Table(models.TableTodo).FindById(&todo, tid)
	if err != nil {
		return result.ServerError(err)
	}
	todo.Position = c.Position
	if c.Status.Set {
		todo.Status = c.Status.Value
	}
	err = orm.Table(models.TableTodo).Update(todo)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(todo)
}

func findTodo(tid int64) (*models.Todo, *result.ApiResult) {
	var t models.Todo
	err := orm.Table(models.TableTodo).FindById(&t, tid)
	if err != nil {
		return nil, result.ServerError(err)
	}
	if t.Id == int64(0) {
		return nil, result.BadRequest("requested todo not found")
	}
	return &t, nil
}

func checkFolderAuthority(uid int64, fid int64) *result.ApiResult {
	var m models.WorkspaceMember
	err := models.WorkspaceMemberQuery(orm.Engine).FindByUserIdAndFolderId(&m, uid, fid)
	if err != nil {
		return result.ServerError(err)
	}
	if m.Id == int64(0) {
		return result.BadRequest("user does not have permission to access the todo")
	}
	return nil
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
