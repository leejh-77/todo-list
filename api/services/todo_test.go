package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
	"todo-list/models"
	"todo-list/optional"
)

func TestGetTodo(t *testing.T) {
	models.ClearTables()

	var (
		u = models.TestUser()
		todo = models.TestTodo()
	)

	ret := GetTodos(u.Id, todo.FolderId)

	assert.Equal(t, http.StatusOK, ret.StatusCode)

	todos := ret.Result.([]models.Todo)
	assert.Equal(t, 1, len(todos))
	assert.Equal(t, todo.Id, todos[0].Id)
}

func TestGetEmptyTodos(t *testing.T) {
	models.ClearTables()

	var (
		u = models.TestUser()
		f = models.TestFolder()
	)

	ret := GetTodos(u.Id, f.Id)

	assert.Equal(t, http.StatusOK, ret.StatusCode)
}

func TestGetTodo_notMember_shouldFail(t *testing.T) {
	var (
		u = models.CreateTestUser("another.user@email.com")
		todo = models.TestTodo()
	)

	ret := GetTodos(u.Id, todo.FolderId)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "user is not a member of the workspace of the folder", ret.Error.Message)
}

func TestCreateTodo(t *testing.T) {
	models.ClearTables()

	var (
		u = models.TestUser()
		f = models.TestFolder()
		c = CreateTodoCommand{
			FolderId:      f.Id,
			Subject:       "test todo",
			Body:          "test todo body",
			Status:        models.TodoStatusNotStarted,
			CompletedTime: 0,
		}
	)

	ret := CreateTodo(u.Id, c)

	assert.Equal(t, http.StatusCreated, ret.StatusCode)

	ret = GetTodos(u.Id, f.Id)

	created := ret.Result.([]models.Todo)[0]
	assert.Equal(t, f.Id, created.FolderId)
	assert.Equal(t, "test todo", created.Subject)
	assert.Equal(t, "test todo body", created.Body)
	assert.Equal(t, models.TodoStatusNotStarted, created.Status)
}

func TestCreateTodo_emptyFolderId_shouldFail(t *testing.T) {
	var (
		u = models.TestUser()
		c = CreateTodoCommand{
			Subject:       "test todo",
			Body:          "test todo body",
			Status:        models.TodoStatusNotStarted,
		}
	)

	ret := CreateTodo(u.Id, c)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "folder id must not be empty", ret.Error.Message)
}

func TestCreateTodo_notMember_shouldFail(t *testing.T) {
	var (
		u = models.CreateTestUser("another.user@email.com")
		f = models.TestFolder()
		c = CreateTodoCommand{
			FolderId:      f.Id,
			Subject:       "test todo",
			Body:          "test todo body",
			Status:        models.TodoStatusNotStarted,
		}
	)

	ret := CreateTodo(u.Id, c)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "user does not have permission to access the todo", ret.Error.Message)
}

func TestCreateTodo_conflictStatusAndCompleteTime_shouldFail(t *testing.T) {
	var (
		u = models.TestUser()
		f = models.TestFolder()
		c = CreateTodoCommand{
			FolderId:      f.Id,
			Subject:       "test todo",
			Body:          "test todo bdoy",
			Status:        models.TodoStatusNotStarted,
			CompletedTime: time.Now().Unix(),
		}
	)

	ret := CreateTodo(u.Id, c)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "uncompleted todo cannot have property `completedTime`", ret.Error.Message)
}

func TestUpdateTodo(t *testing.T) {
	models.ClearTables()

	var (
		u = models.TestUser()
		todo = models.TestTodo()
		c = UpdateTodoCommand{
			Id: todo.Id,
			Subject:       optional.NewString("updated todo subject"),
		}
	)
	ret := UpdateTodo(u.Id, c)
	assert.Equal(t, http.StatusOK, ret.StatusCode)

	ret = GetTodos(u.Id, todo.FolderId)
	todos := ret.Result.([]models.Todo)
	updated := todos[0]

	assert.Equal(t, "updated todo subject", updated.Subject)
	assert.Equal(t, todo.Body, updated.Body)
}

func TestUpdateTodo_emptyId_shouldFail(t *testing.T) {
	var (
		u = models.TestUser()
		c = UpdateTodoCommand{
			Subject:       optional.NewString("updated todo subject"),
		}
	)

	ret := UpdateTodo(u.Id, c)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "id must not be empty", ret.Error.Message)
}

func TestUpdateTodo_invalidId_shouldFail(t *testing.T) {
	var (
		u = models.TestUser()
		c = UpdateTodoCommand{
			Id:            -1,
			Subject: optional.NewString("updated todo subject"),
		}
	)

	ret := UpdateTodo(u.Id, c)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "requested todo not found", ret.Error.Message)
}

func TestUpdateTodo_notMember_shouldFail(t *testing.T) {
	var (
		u = models.CreateTestUser("another.user@email.com")
		todo = models.TestTodo()
		c = UpdateTodoCommand{
			Id:            todo.Id,
			Subject:       optional.NewString("updated todo subject"),
		}
	)

	ret := UpdateTodo(u.Id, c)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "user does not have permission to access the todo", ret.Error.Message)
}

func TestUpdateTodo_conflictStatusAndCompleteTime_shouldFail(t *testing.T) {
	var (
		u = models.TestUser()
		todo = models.TestTodo()
		c = UpdateTodoCommand{
			Id:            todo.Id,
			Status:        optional.NewInt(models.TodoStatusNotStarted),
			CompletedTime: optional.NewInt64(time.Now().Unix()),
		}
	)

	ret := UpdateTodo(u.Id, c)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "uncompleted todo cannot have property `completedTime`", ret.Error.Message)
}

func TestDeleteTodo(t *testing.T) {
	models.ClearTables()

	var (
		u = models.TestUser()
		todo = models.TestTodo()
	)

	ret := DeleteTodo(u.Id, todo.Id)
	assert.Equal(t, http.StatusOK, ret.StatusCode)

	ret = GetTodos(u.Id, todo.FolderId)
	todos := ret.Result.([]models.Todo)
	assert.Equal(t, 0, len(todos))
}

func TestDeleteTodo_emptyId_shouldFail(t *testing.T) {
	var (
		u = models.TestUser()
	)

	ret := DeleteTodo(u.Id, int64(0))

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "id must not be empty", ret.Error.Message)
}

func TestDeleteTodo_invalidId_shouldFail(t *testing.T) {
	var (
		u = models.TestUser()
	)

	ret := DeleteTodo(u.Id, int64(-1))

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "requested todo not found", ret.Error.Message)
}

func TestDeleteTodo_notMember_shouldFail(t *testing.T) {
	var (
		u = models.CreateTestUser("another.user@email.com")
		todo = models.TestTodo()
	)
	ret := DeleteTodo(u.Id, todo.Id)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "user does not have permission to access the todo", ret.Error.Message)
}

