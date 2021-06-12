package models

import "todo-list/orm"

const (
	TodoStatusNotStarted = 0
	TodoStatusInProgress = 1
	TodoStatusCompleted = 2
)

type Todo struct {
	Id int64
	FolderId int64
	UserId int64
	Subject string
	Body string
	Status int
	CompletedTime int64
	Position int
}

type todoTable struct {
	*orm.Table
}
