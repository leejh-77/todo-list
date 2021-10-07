package models

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"todo-list/base"
	"todo-list/orm"
)



func TestMain(m *testing.M) {
	orm.Init(base.TestDBConfig)
	RegisterTables()
	os.Exit(m.Run())
}

func TestFindAll(t *testing.T) {
	_ = orm.Table(TableTodo).DeleteAll()

	for i := 0; i < 3; i++ {
		createTestTodo()
	}

	var arr []Todo
	err := orm.Table(TableTodo).FindAll(&arr)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, len(arr))
}

func TestCreateTodo(t *testing.T) {
	todo := createTestTodo()
	assert.NotNil(t, todo)
}

func TestUpdateTodo(t *testing.T) {
	todo := createTestTodo()
	todo.Subject = "Updated subject"

	err := orm.Table(TableTodo).Update(todo)
	err = orm.Table(TableTodo).FindById(todo, todo.Id)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Updated subject", todo.Subject)
}

func TestDeleteTodo(t *testing.T) {
	todo := createTestTodo()
	err := orm.Table(TableTodo).DeleteById(todo.Id)
	if err != nil {
		t.Fatal(err)
	}
	err = orm.Table(TableTodo).FindById(todo, todo.Id)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(0), todo.Id)
}

func makeTodo(subject string) *Todo {
	todo := new(Todo)
	todo.UserId = int64(1)
	todo.Subject = subject
	todo.Body = "Test Todo Body"
	todo.Status = TodoStatusNotStarted
	return todo
}

func createTestTodo() *Todo {
	todo := makeTodo("Test Todo")
	id, err := orm.Table(TableTodo).Insert(todo)
	if err != nil {
		return nil
	}
	err = orm.Table(TableTodo).FindById(todo, id)
	if err != nil {
		return nil
	}
	return todo
}
