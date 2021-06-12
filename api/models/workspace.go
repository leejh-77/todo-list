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
