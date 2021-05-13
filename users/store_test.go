package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo-list/model"
)

func TestCreateUser(t *testing.T) {
	user := new(model.User)
	user.EmailAddress = "jonghoon.lee@gmail.com"
	user.Password = "passwod@!!"
	user.Username = "Jonghoon Lee"
	user.RegisteredTime = time.Now().Unix()

	err := createUser(user)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, int64(0), user.Id)
}

