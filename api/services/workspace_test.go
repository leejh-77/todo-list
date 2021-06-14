package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"todo-list/models"
)


func TestCreateWorkspace(t *testing.T) {
	clearTables()

	uid := models.TestUser().Id
	c := CreateWorkspaceCommand{
		Name: "Test Workspace",
	}
	ret := CreateWorkspace(uid, c)

	assert.Equal(t, 201, ret.StatusCode)
	ret = GetWorkspaces(uid)

	workspaces := ret.Result.([]models.Workspace)
	assert.NotEqual(t, 0, len(workspaces))
}

func TestCreateWorkspace_invalidName_shouldFail(t *testing.T) {
	c := CreateWorkspaceCommand{
		Name: "",
	}
	ret := CreateWorkspace(models.TestUser().Id, c)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "name must not be empty", ret.Error.Message)
}

func TestDeleteWorkspace(t *testing.T) {
	clearTables()

	u := models.TestUser()
	w := models.TestWorkspace()
	DeleteWorkspace(u.Id, w.Id)

	ret := GetWorkspaces(u.Id)
	ws := ret.Result.([]models.Workspace)
	assert.Equal(t, 0, len(ws))
}