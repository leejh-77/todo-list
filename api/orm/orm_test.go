package orm

import (
	"testing"
	"time"
	"todo-list/models"
)

func TestCreate(t *testing.T) {
	table := NewTable("users", models.User{})
	user := models.User{
		EmailAddress:   "jonghoon.lee@gmail.com",
		Password:       "password#@$!",
		Username:       "Jonghoon Lee",
		RegisteredTime: time.Now().Unix(),
	}
	id, err := table.Save(user)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}

func TestFind(t *testing.T) {
	table := NewTable("users", models.User{})
	arr, err := table.FindAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(arr)
}


