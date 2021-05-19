package auth

import (
	"database/sql"
	"todo-list/base"
	"todo-list/model"
)

func createUser(user *model.User) (int64, error) {
	ret, err := base.DB.Exec(
		"INSERT INTO users (`emailAddress`, `password`, `username`, `registeredTime`) VALUES (?, ?, ?, ?)",
		user.EmailAddress, user.Password, user.Username, user.RegisteredTime)
	if err != nil {
		return -1, err
	}
	return ret.LastInsertId()
}

func findUserById(id int64) (*model.User, error) {
	return scanUser(base.DB.Query("SELECT * FROM users WHERE id = ?", id))
}

func findUserByEmailAddress(email string) (*model.User, error) {
	return scanUser(base.DB.Query("SELECT * FROM users WHERE emailAddress = ?", email))
}

func scanUser(rows *sql.Rows, err error) (*model.User, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	user := new(model.User)
	err = rows.Scan(&user.Id, &user.EmailAddress, &user.Password, &user.Username, &user.RegisteredTime)
	if err != nil {
		return nil, err
	}
	return user, nil
}