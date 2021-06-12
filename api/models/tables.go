package models

import (
	"todo-list/base"
	"todo-list/orm"
)

const (
	tUser = "users"
	tTodo = "todos"
	tWorkspace = "workspaces"
	tWorkspaceMembers = "workspaceMembers"
)

func init() {
	orm.Init(base.DBConfig)

	orm.Register(tUser, User{})
	orm.Register(tTodo, Todo{})
	orm.Register(tWorkspace, Workspace{})
	orm.Register(tWorkspaceMembers, WorkspaceMember{})
}