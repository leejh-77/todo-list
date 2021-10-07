package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"todo-list/models"
)

func TestCreateFolder(t *testing.T) {
	models.ClearTables()

	var (
		w = models.TestWorkspace()
		c = CreateFolderCommand{
			WorkspaceId: w.Id,
			Name:        "test.folder",
		}
	)

	ret := CreateFolder(models.TestUser().Id, c)
	assert.Equal(t, http.StatusCreated, ret.StatusCode)
}

func TestCreateFolder_notMember_shouldFail(t *testing.T) {
	models.ClearTables()

	var (
		w = models.TestWorkspace()
		u = models.CreateTestUser("another.user@email.com")
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
	models.ClearTables()

	var (
		u = models.TestUser()
		f = models.TestFolder()
		w = models.TestWorkspace()
	)

	DeleteFolder(u.Id, f.Id)
	ret := GetFolders(u.Id, w.Id)
	fs := ret.Result.([]models.Folder)

	assert.Equal(t, 0, len(fs))
}

func TestDeleteFolder_notMember_shouldFail(t *testing.T) {
	models.ClearTables()

	var (
		f = models.TestFolder()
		u = models.CreateTestUser("another.user@email.com")
	)

	ret := DeleteFolder(u.Id, f.Id)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "user is not a member of the workspace", ret.Error.Message)
}


