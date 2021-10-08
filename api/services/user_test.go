package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"todo-list/models"
)

func TestGetUser(t *testing.T) {
	u := models.TestUser()
	email := u.EmailAddress

	ret := GetUser(u.Id)

	assert.Equal(t, http.StatusOK, ret.StatusCode)

	res := ret.Result.(*GetUserResponse)
	assert.Equal(t, res.EmailAddress, email)
}
