package models

import (
	"todo-list/base"
	"todo-list/orm"
)

var Users userTable
var Todos todoTable
var Workspaces workspaceTable
var WorkspaceMembers workspaceMemberTable

func init() {
	orm.Init(base.DBConfig)

	Users = userTable{orm.NewTable("users", User{})}
	Todos = todoTable{orm.NewTable("todos", Todo{})}
	Workspaces = workspaceTable{orm.NewTable("workspaces", Workspace{})}
	WorkspaceMembers = workspaceMemberTable{orm.NewTable("workspaceMembers", WorkspaceMember{})}
}