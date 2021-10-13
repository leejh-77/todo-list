package models

import "todo-list/orm"

const (
	TodoStatusNotStarted = 0
	TodoStatusInProgress = 1
	TodoStatusCompleted = 2
)

type Todo struct {
	Id int64 `json:"id"`
	FolderId int64 `json:"folderId"`
	UserId int64 `json:"userId"`
	Subject string `json:"subject"`
	Body string `json:"body"`
	Status int `json:"status"`
	CompletedTime int64 `json:"completedTime"`
	Position float32 `json:"position"`
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

func (q *todoQuery) FindLastPosition(fid int64, status int) (float32, error) {
	var t Todo
	err := q.s.Table(TableTodo).Find(&t, "folderId = ? AND status = ? ORDER BY position DESC LIMIT 1", fid, status)
	if err != nil {
		return .0, err
	}
	if t.Id == int64(0) {
		return .0, nil
	}
	return t.Position, nil
}

