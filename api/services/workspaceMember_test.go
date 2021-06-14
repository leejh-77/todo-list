package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"todo-list/models"
)

func TestGetWorkspaceMembers(t *testing.T) {
	u := models.TestUser()

	w := models.TestWorkspace()
	ret := GetWorkspaceMembers(u.Id, w.Id)

	data := ret.Result.([]GetUserResponseData)
	assert.Equal(t, 1, len(data))

	d := data[0]
	assert.Equal(t, u.Username, d.Name)
	assert.Equal(t, u.EmailAddress, d.EmailAddress)
	assert.True(t, d.IsOwner)
}

func TestDeleteWorkspaceMember(t *testing.T) {
	u := models.TestUser()
	w := models.TestWorkspace()

	ret := DeleteWorkspaceMember(u.Id, w.Id)

	assert.Equal(t, 200, ret.StatusCode)

	ret = GetWorkspaces(u.Id)
	ws := ret.Result.([]models.Workspace)

	assert.Equal(t, 0, len(ws))
}

func TestDeleteWorkspaceMember_checkDeleteWhenZeroMembers(t *testing.T) {
	u, w := models.TestUser(), models.TestWorkspace()

	DeleteWorkspaceMember(u.Id, w.Id)

	var found models.Workspace
	err := models.Workspaces.FindById(&found, w.Id)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(0), found.Id)
}
