package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"todo-list/models"
	"todo-list/orm"
	"todo-list/test"
)


func TestCreateWorkspace(t *testing.T) {
	test.ClearTables()

	uid := test.TestUser().Id
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
	ret := CreateWorkspace(test.TestUser().Id, c)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "name must not be empty", ret.Error.Message)
}

func TestDeleteWorkspace(t *testing.T) {
	test.ClearTables()

	u, w := test.TestUser(), test.TestWorkspace()
	DeleteWorkspace(u.Id, w.Id)

	ret := GetWorkspaces(u.Id)
	ws := ret.Result.([]models.Workspace)
	assert.Equal(t, 0, len(ws))
}

func TestDeleteWorkspace_notMember_shouldFail(t *testing.T) {
	w := test.TestWorkspace()
	u := test.CreateTestUser("another.user@email.com")

	ret := DeleteWorkspace(u.Id, w.Id)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "user is not a member of the workspace", ret.Error.Message)
}

func TestDeleteWorkspace_permissionDenied_shouldFail(t *testing.T) {
	test.ClearTables()

	w := test.TestWorkspace()
	u := test.CreateTestUser("another.user@email.com")

	ret := AddWorkspaceMember(u.Id, w.Id)

	assert.Equal(t, http.StatusOK, ret.StatusCode)

	ret = DeleteWorkspace(u.Id, w.Id)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "user does not have permission to delete workspace", ret.Error.Message)
}

func TestGetMembers(t *testing.T) {
	test.ClearTables()
	var (
		u = test.TestUser()
		w = test.TestWorkspace()
	)
	ret := GetWorkspaceMembers(u.Id, w.Id)

	data := ret.Result.([]GetMemberResponseData)
	assert.Equal(t, 1, len(data))

	d := data[0]
	assert.Equal(t, u.Username, d.Name)
	assert.Equal(t, u.EmailAddress, d.EmailAddress)
	assert.True(t, d.IsOwner)
}

func TestDeleteMember(t *testing.T) {
	u := test.TestUser()
	w := test.TestWorkspace()

	ret := DeleteWorkspaceMember(u.Id, w.Id)

	assert.Equal(t, 200, ret.StatusCode)

	ret = GetWorkspaces(u.Id)
	ws := ret.Result.([]models.Workspace)

	assert.Equal(t, 0, len(ws))
}

func TestDeleteMember_notExist_shouldFail(t *testing.T) {
	w := test.TestWorkspace()

	ret := DeleteWorkspaceMember(w.Id, int64(93874)) // 존재하지 않는 user id

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "user is not a member of the workspace", ret.Error.Message)
}

func TestDeleteMember_checkDeleteWorkspaceWhenZeroMembers(t *testing.T) {
	u, w := test.TestUser(), test.TestWorkspace()

	DeleteWorkspaceMember(u.Id, w.Id)

	var found models.Workspace
	err := orm.Table(models.TableWorkspace).FindById(&found, w.Id)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(0), found.Id)
}
