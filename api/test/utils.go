package test

import (
	"log"
	"strconv"
	"time"
	"todo-list/base"
)

func BeforeTest() {
	_, err := base.DB.Exec("TRUNCATE TABLE users")
	_, err = base.DB.Exec("TRUNCATE TABLE todos")
	if err != nil {
		log.Fatal(err)
	}
}

func UniqueString(str string) string {
	return str + strconv.FormatInt(time.Now().UnixNano(), 10)
}

