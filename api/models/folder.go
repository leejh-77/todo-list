package models

import "todo-list/orm"

type Folder struct {
	Id int64
	Name string
	WorkspaceId int64
}

type folderTable struct {
	*orm.ORMTable
}