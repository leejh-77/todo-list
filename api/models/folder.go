package models

type Folder struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	WorkspaceId int64 `json:"workspaceId"`
}
