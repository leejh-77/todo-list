package orm

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo-list/models"
)

func init() {
	models.BeforeTest()
}

func TestCreate(t *testing.T) {
	table := NewTable("users", models.User{})
	user := userMock()
	id, err := table.Insert(user)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, id, int64(-1))
}

func TestUpdate(t *testing.T) {
	table := NewTable("users", models.User{})
	user := userMock()
	id, err := table.Insert(user)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, id, int64(-1))

	u, err := table.FindById(id)
	if err != nil {
		t.Fatal(err)
	}
	user = u.(*models.User)
	user.Username = "Todo-list"
	err = table.Update(user)
	if err != nil {
		t.Fatal(err)
	}
	u, err = table.FindById(id)
	if err != nil {
		t.Fatal(err)
	}
	user = u.(*models.User)
	assert.Equal(t, user.Username, "Todo-list")
}

func TestDelete(t *testing.T) {

}

func TestFind(t *testing.T) {
	table := NewTable("users", models.User{})
	arr, err := table.FindAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(arr)
}

func userMock() *models.User {
	return &models.User{
		EmailAddress:   "jonghoon.lee@gmail.com",
		Password:       "password#@$!",
		Username:       "Jonghoon Lee",
		RegisteredTime: time.Now().Unix(),
	}
}

