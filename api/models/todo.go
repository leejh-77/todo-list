package models

const (
	TodoStatusNotStarted = 0
	TodoStatusInProgress = 1
	TodoStatusCompleted = 2
)

type Todo struct {
	Id int64
	UserId int64
	Subject string
	Body string
	Status int
	CompletedTime int64
}

func SaveTodo(t *Todo) (int64, error) {
	if t.Id == 0 {
		return createTodo(t)
	}
	return t.Id, updateTodo(t)
}

func DeleteTodo(id int64) error {
	_, err := DB.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func createTodo(t *Todo) (int64, error) {
	ret, err := DB.Exec(
		"INSERT INTO todos (userId, subject, body, status, completedTime) VALUES (?, ?, ?, ?, ?)",
		t.UserId, t.Subject, t.Body, t.Status, t.CompletedTime)
	if err != nil {
		return -1, err
	}
	return ret.LastInsertId()
}

func updateTodo(t *Todo) error {
	_, err := DB.Exec(
		"UPDATE todos SET userId = ?, subject = ?, body = ?, status = ?, completedTime = ? WHERE id = ?",
		t.UserId, t.Subject, t.Body, t.Status, t.CompletedTime, t.Id)
	if err != nil {
		return err
	}
	return nil
}

func FindById(id int64) (*Todo, error) {
	return singleResult("id = ?", id)
}

func FindAll() ([]*Todo, error) {
	return multiResult("")
}

func FindByUserId(id int64) ([]*Todo, error) {
	return multiResult("userId = ?", id)
}

func singleResult(where string, args... interface{}) (*Todo, error) {
	arr, err := multiResult(where, args...)
	if err != nil {
		return nil, err
	}
	if len(arr) == 0 {
		return nil, nil
	}
	return arr[0], nil
}

func multiResult(where string, args... interface{}) ([]*Todo, error) {
	query := "SELECT * FROM todos"
	if len(where) > 0 {
		query = query + " WHERE " + where
	}
	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arr := make([]*Todo, 0)
	for rows.Next() {
		todo := new(Todo)
		err = rows.Scan(&todo.Id, &todo.UserId, &todo.Subject, &todo.Body, &todo.Status, &todo.CompletedTime)
		arr = append(arr, todo)
	}
	return arr, nil
}
