package models

import (
	"todo-list/orm"
)

const (
	TableUser            = "users"
	TableTodo            = "todos"
	TableWorkspace       = "workspaces"
	TableWorkspaceMember = "workspaceMembers"
	TableFolder          = "folders"
)

func RegisterTables() {
	orm.Register(TableUser, User{})
	orm.Register(TableTodo, Todo{})
	orm.Register(TableWorkspace, Workspace{})
	orm.Register(TableWorkspaceMember, WorkspaceMember{})
	orm.Register(TableFolder, Folder{})
}