package models

import "todo-list/orm"

type Workspace struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	CreatedTime int64 `json:"createdTime"`
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

func (q *workspaceQuery) FindByNameLike(ws *[]Workspace, uid int64, name string) error {
	name = "%" + name + "%"
	return q.s.Table(TableWorkspace).Find(ws, "name LIKE ? AND id NOT IN (SELECT workspaceId FROM " +
		TableWorkspaceMember + " WHERE userId = ?)", name, uid)
}
