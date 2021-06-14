package models

import (
	"log"
	"time"
)

func TestUser() *User {
	email := "todo.test.user@gmail.com"

	var u User
	_ = Users.FindByEmailAddress(&u, email)
	if u.Id > 0 {
		return &u
	}
	u = User{
		EmailAddress:   email,
		Password:       "password!@#$",
		Username:       "Jonghoon Lee",
		RegisteredTime: time.Now().Unix(),
	}
	id, err := Users.Insert(&u)
	if err != nil {
		log.Fatal(err)
	}
	u.Id = id
	return &u
}

func TestWorkspace() *Workspace {
	name := "test workspace"
	user := TestUser()

	var w Workspace
	err := Workspaces.Find(&w,
		"name = ? AND id IN (SELECT workspaceId FROM workspaceMembers WHERE userId = ?)",
		user.Id,
		name)
	if err != nil {
		log.Fatal(err)
	}
	if w.Id != int64(0) {
		return &w
	}

	w = Workspace{
		Name:        name,
		CreatedTime: time.Now().Unix(),
	}
	id, err := Workspaces.Insert(&w)
	if err != nil {
		log.Fatal(err)
	}
	w.Id = id
	return &w
}
