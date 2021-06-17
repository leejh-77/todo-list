package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"todo-list/models"
	"todo-list/test"
)

func TestCreateFolder(t *testing.T) {
	test.ClearTables()

	var (
		w = test.TestWorkspace()
		c = CreateFolderCommand{
			WorkspaceId: w.Id,
			Name:        "test.folder",
		}
	)

	ret := CreateFolder(test.TestUser().Id, c)
	assert.Equal(t, http.StatusCreated, ret.StatusCode)
}

func TestCreateFolder_notMember_shouldFaile(t *testing.T) {
	test.ClearTables()

	var (
		w = test.TestWorkspace()
		u = test.CreateTestUser("another.user@email.com")
		c = CreateFolderCommand{
			WorkspaceId: w.Id,
			Name: "test.folder",
		}
	)

	ret := CreateFolder(u.Id, c)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "user is not a member of the workspace", ret.Error.Message)
}

func TestDeleteFolder(t *testing.T) {
	test.ClearTables()

	var (
		u = test.TestUser()
		f = test.TestFolder()
		w = test.TestWorkspace()
	)

	DeleteFolder(u.Id, f.Id)
	ret := GetFolders(u.Id, w.Id)
	fs := ret.Result.([]models.Folder)

	assert.Equal(t, 0, len(fs))
}

func TestDeleteFolder_notMember_shouldFail(t *testing.T) {
	test.ClearTables()

	var (
		f = test.TestFolder()
		u = test.CreateTestUser("another.user@email.com")
	)

	ret := DeleteFolder(u.Id, f.Id)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "user is not a member of the workspace", ret.Error.Message)
}


