package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	w := TestWorkspace()
	f := Folder{
		Name:        "test folder",
		WorkspaceId: w.Id,
	}

	id, err := Folders.Insert(&f)
	if err != nil {
		t.Error(err)
	}
	assert.Greater(t, id, int64(0))
}