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

type todoQuery struct {
	s orm.Session
}

func TodoQuery(s orm.Session) *todoQuery {
	return &todoQuery{s}
}

func (q *todoQuery) FindByFolderId(ts *[]Todo, fid int64) error {
	return q.s.Table(TableTodo).Find(ts, "folderId = ?", fid)
}
