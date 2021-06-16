package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"todo-list/models"
	"todo-list/orm"
)

func TestGetMembers(t *testing.T) {
	var (
		u = TestUser()
		w = TestWorkspace()
	)
	ret := GetWorkspaceMembers(u.Id, w.Id)

	data := ret.Result.([]GetUserResponseData)
	assert.Equal(t, 1, len(data))

	d := data[0]
	assert.Equal(t, u.Username, d.Name)
	assert.Equal(t, u.EmailAddress, d.EmailAddress)
	assert.True(t, d.IsOwner)
}

func TestDeleteMember(t *testing.T) {
	u := TestUser()
	w := TestWorkspace()

	ret := DeleteWorkspaceMember(u.Id, w.Id)

	assert.Equal(t, 200, ret.StatusCode)

	ret = GetWorkspaces(u.Id)
	ws := ret.Result.([]models.Workspace)

	assert.Equal(t, 0, len(ws))
}

func TestDeleteMember_notExist_shouldFail(t *testing.T) {
	w := TestWorkspace()

	ret := DeleteWorkspaceMember(w.Id, int64(93874)) // 존재하지 않는 user id

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "user is not a member of the workspace", ret.Error.Message)
}

func TestDeleteMember_checkDeleteWorkspaceWhenZeroMembers(t *testing.T) {
	u, w := TestUser(), TestWorkspace()

	DeleteWorkspaceMember(u.Id, w.Id)

	var found models.Workspace
	err := orm.Table(models.TableWorkspace).FindById(&found, w.Id)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(0), found.Id)
}
