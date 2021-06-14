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

func (t *folderTable) FindByWorkspaceId(fs *[]Folder, wid int64) error {
	return t.Find(fs, "workspaceId = ?", wid)
}