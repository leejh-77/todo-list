package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"todo-list/models"
)


func TestCreateWorkspace(t *testing.T) {
	_ = models.Workspaces.DeleteAll()
	_ = models.WorkspaceMembers.DeleteAll()

	u := testUser()
	e := createAuthorizedContext(u)

	c := CreateWorkspaceCommand{
		Name: "Test Workspace",
	}
	e.Set("command", c)
	ret := CreateWorkspace(e)

	assert.Equal(t, 201, ret.StatusCode)

	var member models.WorkspaceMember
	err := models.WorkspaceMembers.Find(&member, "userId = ?", u.Id)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, member)
	assert.Equal(t, member.Type, models.MemberTypeOwner)
}

func TestCreateWorkspace_invalidName(t *testing.T) {
	e := createAuthorizedContext(testUser())

	c := CreateWorkspaceCommand{
		Name: "",
	}
	e.Set("command", c)
	ret := CreateWorkspace(e)

	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
	assert.Equal(t, "name must not be empty", ret.Error.Message)
}

