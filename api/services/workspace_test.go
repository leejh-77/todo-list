package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateWorkspace(t *testing.T) {
	

	e := createDummyContext()
	e.Set("command", e)
	ret := CreateWorkspace(e)

	assert.Equal(t, 201, ret.StatusCode)
}

