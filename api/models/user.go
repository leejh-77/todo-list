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
	*orm.Table
}

var Users = userTable{
	orm.NewTable("users", User{}),
}

func (t *userTable) FindByEmailAddress(u *User, email string) error {
	return t.Table.FindOne(u, "emailAddress = ?", email)
}