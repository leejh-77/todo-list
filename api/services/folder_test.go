package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"todo-list/models"
)

func TestCreateFolder(t *testing.T) {
	clearTables()

	w := TestWorkspace()
	c := CreateFolderCommand{
		WorkspaceId: w.Id,
		Name:        "test.folder",
	}
	ret := CreateFolder(TestUser().Id, c)

	assert.Equal(t, http.StatusCreated, ret.StatusCode)
}

func TestDeleteFolder(t *testing.T) {
	clearTables()

	u := TestUser()
	f := TestFolder()

	DeleteFolder(u.Id, f.Id)

	ret := GetFolders(u.Id, TestWorkspace().Id)
	fs := ret.Result.([]models.Folder)

	assert.Equal(t, 0, len(fs))
}