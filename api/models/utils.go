package models

import "log"

func BeforeTest() {
	_, err := DB.Exec("TRUNCATE TABLE users")
	_, err = DB.Exec("TRUNCATE TABLE todos")
	if err != nil {
		log.Fatal(err)
	}
}

