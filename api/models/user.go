package models

import "todo-list/orm"

type User struct {
	Id int64
	EmailAddress string
	Password string
	Username string
	RegisteredTime int64
}

type userQuery struct {
	Session orm.Session
}

func UserQuery(s orm.Session) *userQuery {
	return &userQuery{Session: s}
}

func (q *userQuery) FindByEmailAddress(u *User, email string) error {
	return q.Session.Table(TableUser).Find(u, "emailAddress = ?", email)
}

func (q *userQuery) FindByWorkspace(u *[]User, wid int64) error {
	return q.Session.Table(TableUser).Find(u, "id IN (SELECT userId FROM workspaceMembers WHERE workspaceId = ?)", wid)
}


