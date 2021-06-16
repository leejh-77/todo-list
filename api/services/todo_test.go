package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"todo-list/models"
)

func TestGetTodo(t *testing.T) {
	clearTables()

	var (
		u = TestUser()
		todo = TestTodo()
	)

	ret := GetTodo(u.Id, todo.FolderId)

	assert.Equal(t, http.StatusOK, ret.StatusCode)

	todos := ret.Result.([]models.Todo)
	assert.Equal(t, 1, len(todos))
	assert.Equal(t, todo.Id, todos[0].Id)
}

func TestGetTodo_notMember_shouldFail(t *testing.T) {
	clearTables()

	var (
		u = createTestUser("another.user@email.com")
		todo = TestTodo()
	)

	ret := GetTodo(u.Id, todo.FolderId)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "user is not a member of the workspace of the folder", ret.Error.Message)
}

func TestCreateTodo(t *testing.T) {
	clearTables()

	var (
		u = TestUser()
		f = TestFolder()
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

	ret = GetTodo(u.Id, f.Id)

	created := ret.Result.([]models.Todo)[0]
	assert.Equal(t, f.Id, created.FolderId)
	assert.Equal(t, "test todo", created.Subject)
	assert.Equal(t, "test todo body", created.Body)
	assert.Equal(t, models.TodoStatusNotStarted, created.Status)
}

func TestCreateTodo_emptyFolderId_shouldFail(t *testing.T) {

}

func TestCreateTodo_notMember_shouldFail(t *testing.T) {

}

func TestCreateTodo_conflictStatusAndCompleteTime(t *testing.T) {

}

func TestUpdateTodo(t *testing.T) {

}

func TestUpdateTodo_notMember_shouldFail(t *testing.T) {

}

func TestUpdateTodo_conflictStatusAndCompleteTime(t *testing.T) {

}

func TestDeleteTodo(t *testing.T) {

}

func TestDeleteTodo_notMember_shouldFail(t *testing.T) {

}

func TestChangePosition(t *testing.T) {

}
