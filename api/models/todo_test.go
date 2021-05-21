package models

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"todo-list/base"
	"todo-list/test"
)

func init() {
	test.BeforeTest()
}

func TestFindAll(t *testing.T) {
	deleteAll()

	for i := 0; i < 3; i++ {
		_, err := createTestTodo("test todo")
		if err != nil {
			t.Fatal(err)
		}
	}
	todos, err := FindAll()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, len(todos))
}

func TestCreate(t *testing.T) {
	id, err := createTestTodo("Test Todo")
	if err != nil {
		t.Fatal(err)
	}
	assert.Greater(t, id, int64(0))
}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}

func TestComplete(t *testing.T) {

}

func makeTodo(subject string) *Todo {
	todo := new(Todo)
	todo.UserId = 1
	todo.Subject = subject
	todo.Body = "Test Todo Body"
	todo.Status = TodoStatusNotStarted
	return todo
}

func createTestTodo(subject string) (int64, error) {
	todo := makeTodo(subject)
	return CreateTodo(todo)
}

func deleteAll() {
	_, err := base.DB.Exec("TRUNCATE TABLE todos")
	if err != nil {
		log.Fatal(err)
	}
}