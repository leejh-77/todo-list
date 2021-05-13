package users

import (
	"database/sql"
	"todo-list/model"
	"todo-list/store"
)

func createUser(user *model.User) error {
	_, err := store.DB.Exec(
		"INSERT INTO users (`emailAddress`, `password`, `username`, `registeredTime`) VALUES (?, ?, ?, ?)",
		user.EmailAddress, user.Password, user.Username, user.RegisteredTime)
	if err != nil {
		return err
	}
	return nil
}

func findUserById(id int64) (*model.User, error) {
	return scanUser(store.DB.Query("SELECT * FROM users WHERE id = ?", id))
}


func findUserByEmailAddress(email string) (*model.User, error) {
	return scanUser(store.DB.Query("SELECT * FROM users WHERE emailAddress = ?", email))
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