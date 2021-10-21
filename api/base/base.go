package base

import "todo-list/orm"

const JWTSecret = "todo-list-jwt-secret"

var DBConfig = orm.DatabaseConfig{
	Driver:   "mysql",
	Host:     "localhost:3306",
	User:     "root",
	Password: "1234",
	Name:     "todo",
}

var TestDBConfig = orm.DatabaseConfig{
	Driver:   "mysql",
	Host:     "localhost:3306",
	User:     "root",
	Password: "1234",
	Name:     "todo_test",
}

