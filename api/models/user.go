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

func UserTable() *userTable {
	return &userTable{
		orm.Table(tUser),
	}
}

func (t *userTable) FindByEmailAddress(u *User, email string) error {
	return t.FindOne(u, "emailAddress = ?", email)
}