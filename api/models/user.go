package models

import (
	"todo-list/orm"
)

type User struct {
	Id int64
	EmailAddress string
	Password string
	Username string
	RegisteredTime int64
}

type userTable struct {
	*orm.ORMTable
}

func (t *userTable) FindByEmailAddress(u *User, email string) error {
	return t.FindOne(u, "emailAddress = ?", email)
}

func (t *userTable) FindByWorkspace(us *[]User, wid int64) error {
	return t.Find(us, "id IN (SELECT userId FROM workspaceMembers WHERE workspaceId = ?)", wid)
}
