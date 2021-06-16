package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"todo-list/models"
)

func TestCreateFolder(t *testing.T) {
	clearTables()

	var (
		w = TestWorkspace()
		c = CreateFolderCommand{
			WorkspaceId: w.Id,
			Name:        "test.folder",
		}
	)

	ret := CreateFolder(TestUser().Id, c)
	assert.Equal(t, http.StatusCreated, ret.StatusCode)
}

func TestCreateFolder_notMember_shouldFaile(t *testing.T) {
	clearTables()

	var (
		w = TestWorkspace()
		u = createTestUser("another.user@email.com")
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
	clearTables()

	var (
		u = TestUser()
		f = TestFolder()
		w = TestWorkspace()
	)

	DeleteFolder(u.Id, f.Id)
	ret := GetFolders(u.Id, w.Id)
	fs := ret.Result.([]models.Folder)

	assert.Equal(t, 0, len(fs))
}

func TestDeleteFolder_notMember_shouldFail(t *testing.T) {
	clearTables()

	var (
		f = TestFolder()
		u = createTestUser("another.user@email.com")
	)

	ret := DeleteFolder(u.Id, f.Id)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "user is not a member of the workspace", ret.Error.Message)
}


