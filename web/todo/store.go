package todo

import (
	"todo-list/base"
	"todo-list/model"
)

func createTodo(t *model.Todo) (int64, error) {
	ret, err := base.DB.Exec(
		"INSERT INTO todos (userId, subject, body, status, completedTime) VALUES (?, ?, ?, ?, ?)",
		t.UserId, t.Subject, t.Body, t.Status, t.CompletedTime)
	if err != nil {
		return -1, err
	}
	return ret.LastInsertId()
}

func findById(id int64) (*model.Todo, error) {
	return singleResult("id = ?", id)
}

func findAll() ([]*model.Todo, error) {
	return multiResult("")
}

func findByUserId(id int64) ([]*model.Todo, error) {
	return multiResult("userId = ?", id)
}

func singleResult(where string, args... interface{}) (*model.Todo, error) {
	arr, err := multiResult(where, args)
	if err != nil {
		return nil, err
	}
	return arr[0], nil
}

func multiResult(where string, args... interface{}) ([]*model.Todo, error) {
	query := "SELECT * FROM todos"
	if len(where) > 0 {
		query = query + " WHERE " + where
	}
	rows, err := base.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arr := make([]*model.Todo, 0)
	for rows.Next() {
		todo := new(model.Todo)
		err = rows.Scan(&todo.Id, &todo.UserId, &todo.Subject, &todo.Body, &todo.Status, &todo.CompletedTime)
		arr = append(arr, todo)
	}
	return arr, nil
}
