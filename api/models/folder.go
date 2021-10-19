package models

import "todo-list/orm"

type Folder struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	WorkspaceId int64 `json:"workspaceId"`
}

type folderQuery struct {
	s orm.Session
}

func FolderQuery(s orm.Session) *folderQuery {
	return &folderQuery{s: s}
}

func (q *folderQuery) FindByWorkspaceId(fs *[]Folder, wid int64) error {
	return q.s.Table(TableFolder).Find(fs, "workspaceId = ?", wid)
}