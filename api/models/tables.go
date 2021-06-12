package models

import (
	"todo-list/base"
	"todo-list/orm"
)

const (
	TableUser             = "users"
	TableTodo             = "todos"
	TableWorkspace        = "workspaces"
	TableWorkspaceMembers = "workspaceMembers"
)

var Users *userTable
var Todos *todoTable
var Workspaces *workspaceTable
var WorkspaceMembers *workspaceMemberTable

func init() {
	orm.Init(base.DBConfig)

	orm.Register(TableUser, User{})
	orm.Register(TableTodo, Todo{})
	orm.Register(TableWorkspace, Workspace{})
	orm.Register(TableWorkspaceMembers, WorkspaceMember{})

	Users = &userTable{orm.Table(TableUser)}
	Todos = &todoTable{orm.Table(TableTodo)}
	Workspaces = &workspaceTable{orm.Table(TableWorkspace)}
	WorkspaceMembers = &workspaceMemberTable{orm.Table(TableWorkspaceMembers)}
}