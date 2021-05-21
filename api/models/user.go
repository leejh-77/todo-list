package models

import (
	"database/sql"
	"todo-list/base"
)

type User struct {
	Id int64
	EmailAddress string
	Password string
	Username string
	RegisteredTime int64
}

func CreateUser(user *User) (int64, error) {
	ret, err := base.DB.Exec(
		"INSERT INTO users (emailAddress, password, username, registeredTime) VALUES (?, ?, ?, ?)",
		user.EmailAddress, user.Password, user.Username, user.RegisteredTime)
	if err != nil {
		return -1, err
	}
	return ret.LastInsertId()
}

func FindUserById(id int64) (*User, error) {
	return scanUser(base.DB.Query("SELECT * FROM users WHERE id = ?", id))
}

func FindUserByEmailAddress(email string) (*User, error) {
	return scanUser(base.DB.Query("SELECT * FROM users WHERE emailAddress = ?", email))
}

func scanUser(rows *sql.Rows, err error) (*User, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	user := new(User)
	err = rows.Scan(&user.Id, &user.EmailAddress, &user.Password, &user.Username, &user.RegisteredTime)
	if err != nil {
		return nil, err
	}
	return user, nil
}