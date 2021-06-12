package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateWorkspace(t *testing.T) {
	e := dummyContext()
	e.Set("command", CreateWorkspaceCommand{})
	ret := CreateWorkspace(e)

	assert.Equal(t, 201, ret.StatusCode)
}

