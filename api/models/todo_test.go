package models

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)


func TestFindAll(t *testing.T) {
	for i := 0; i < 3; i++ {
		createTestTodo()
	}

	var arr []Todo
	err := Todos.FindAll(&arr)
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

	err := Todos.Update(todo)
	err = Todos.FindById(todo, todo.Id)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Updated subject", todo.Subject)
}

func TestDeleteTodo(t *testing.T) {
	todo := createTestTodo()
	err := Todos.Delete(todo.Id)
	if err != nil {
		t.Fatal(err)
	}
	err = Todos.FindById(todo, todo.Id)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(0), todo.Id)
}

func makeTodo(subject string) *Todo {
	todo := new(Todo)
	todo.UserId = testUser().Id
	todo.Subject = subject
	todo.Body = "Test Todo Body"
	todo.Status = TodoStatusNotStarted
	return todo
}

func createTestTodo() *Todo {
	todo := makeTodo("Test Todo")
	id, err := Todos.Insert(todo)
	if err != nil {
		return nil
	}
	err = Todos.FindById(todo, id)
	if err != nil {
		return nil
	}
	return todo
}

func testUser() *User {
	email := "todo.test.user@gmail.com"

	var u User
	_ = Users.FindByEmailAddress(&u, email)
	if u.Id > 0 {
		return &u
	}
	u = User{
		EmailAddress:   email,
		Password:       "password!@#$",
		Username:       "Jonghoon Lee",
		RegisteredTime: time.Now().Unix(),
	}
	id, err := Users.Insert(&u)
	if err != nil {
		log.Fatal(err)
	}
	u.Id = id
	return &u
}
