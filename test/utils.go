package test

import (
	"log"
	"todo-list/base"
)

func TruncateTables() {
	_, err := base.DB.Exec("TRUNCATE TABLE users")
	if err != nil {
		log.Fatal(err)
	}
}
