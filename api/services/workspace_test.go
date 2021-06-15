package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"todo-list/models"
)


func TestCreateWorkspace(t *testing.T) {
	clearTables()

	uid := TestUser().Id
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
	ret := CreateWorkspace(TestUser().Id, c)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "name must not be empty", ret.Error.Message)
}

func TestDeleteWorkspace(t *testing.T) {
	clearTables()

	u, w := TestUser(), TestWorkspace()
	DeleteWorkspace(u.Id, w.Id)

	ret := GetWorkspaces(u.Id)
	ws := ret.Result.([]models.Workspace)
	assert.Equal(t, 0, len(ws))
}

func TestDeleteWorkspace_notMember_shouldFail(t *testing.T) {
	w := TestWorkspace()
	u := createTestUser("another.user@email.com")

	ret := DeleteWorkspace(u.Id, w.Id)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "user is not a member of the workspace", ret.Error.Message)
}

func TestDeleteWorkspace_permissionDenied_shouldFail(t *testing.T) {
	//w := TestWorkspace()
	//u := createTestUser("another.user@email.com")

	//ret := DeleteWorkspace(u.Id, w.Id)
}