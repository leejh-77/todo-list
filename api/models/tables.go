package models

import (
	"todo-list/base"
	"todo-list/orm"
)

const (
	TableUser            = "users"
	TableTodo            = "todos"
	TableWorkspace       = "workspaces"
	TableWorkspaceMember = "workspaceMembers"
	TableFolder          = "folders"
)

var Users *userTable
var Todos *todoTable
var Workspaces *workspaceTable
var WorkspaceMembers *workspaceMemberTable
var Folders *folderTable

func init() {
	orm.Init(base.DBConfig)

	orm.Register(TableUser, User{})
	orm.Register(TableTodo, Todo{})
	orm.Register(TableWorkspace, Workspace{})
	orm.Register(TableWorkspaceMember, WorkspaceMember{})
	orm.Register(TableFolder, Folder{})

	Users = &userTable{orm.Table(TableUser)}
	Todos = &todoTable{orm.Table(TableTodo)}
	Workspaces = &workspaceTable{orm.Table(TableWorkspace)}
	WorkspaceMembers = &workspaceMemberTable{orm.Table(TableWorkspaceMember)}
	Folders = &folderTable{orm.Table(TableFolder)}
}