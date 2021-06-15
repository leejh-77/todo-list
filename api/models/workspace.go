package models

import "todo-list/orm"

type Workspace struct {
	Id int64
	Name string
	CreatedTime int64
}

type workspaceQuery struct {
	s orm.Session
}

func WorkspaceQuery(s orm.Session) *workspaceQuery {
	return &workspaceQuery{s}
}

func (q *workspaceQuery) FindByUserId(ws *[]Workspace, uid int64) error {
	return q.s.Table(TableWorkspace).Find(ws, "id IN (SELECT workspaceId FROM workspaceMembers WHERE userId = ?)", uid)
}
