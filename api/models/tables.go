package models

import (
	"todo-list/base"
	"todo-list/orm"
)

var Users userTable
var Todos todoTable

func init() {
	orm.Init(base.DBConfig)

	Users = userTable{orm.NewTable("users", User{})}
	Todos = todoTable{orm.NewTable("todos", Todo{})}
}