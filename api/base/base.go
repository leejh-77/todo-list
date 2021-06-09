package base

import "todo-list/orm"

const JWTSecret = "todo-list-jwt-secret"

var DBConfig = orm.DatabaseConfig{
	Driver:   "mysql",
	Host:     "127.0.0.1:3306",
	User:     "root",
	Password: "1234",
	Name:     "todo",
}
