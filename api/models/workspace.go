package models

import "todo-list/orm"

type Workspace struct {
	Id int64
	Name string
	CreatedTime int64
}

type workspaceTable struct {
	*orm.ORMTable
}

func (t *workspaceTable) FindParticipatingWorkspaces(ws *[]Workspace, uid int64) error {
	return t.Find(ws, "id IN (SELECT userId FROM workspaceMembers WHERE userId = ?)", uid)
}
