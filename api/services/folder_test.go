package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"todo-list/models"
)

func TestCreateFolder(t *testing.T) {
	clearTables()

	w := models.TestWorkspace()
	c := CreateFolderCommand{
		WorkspaceId: w.Id,
		Name:        "test.folder",
	}
	ret := CreateFolder(models.TestUser().Id, c)

	assert.Equal(t, http.StatusCreated, ret.StatusCode)
}

func TestDeleteFolder(t *testing.T) {

}