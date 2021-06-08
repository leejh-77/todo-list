package models

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

var testUser *User

func init() {
	BeforeTest()

	u := &User{
		EmailAddress:   "jonghoon.lee@gmail.com",
		Password:       "password!@#$",
		Username:       "Jonghoon Lee",
		RegisteredTime: time.Now().Unix(),
	}
	id, err := Users.Insert(u)
	if err != nil {
		log.Fatal(err)
	}
	u.Id = id
	testUser = u
}

func TestFindAll(t *testing.T) {
	BeforeTest()

	for i := 0; i < 3; i++ {
		createTestTodo()
	}

	arr := make([]*Todo, 0)
	err := Todos.FindAll(&arr)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, len(arr))
}

func TestCreate(t *testing.T) {
	todo := createTestTodo()
	assert.NotNil(t, todo)
}

func TestUpdate(t *testing.T) {
	todo := createTestTodo()
	todo.Subject = "Updated subject"

	err := Todos.Update(todo)
	err = Todos.FindById(todo, todo.Id)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Updated subject", todo.Subject)
}

func TestDelete(t *testing.T) {
	todo := createTestTodo()
	err := Todos.Delete(todo.Id)
	if err != nil {
		t.Fatal(err)
	}
	err = Todos.FindById(todo, todo.Id)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, todo)
}

func makeTodo(subject string) *Todo {
	todo := new(Todo)
	todo.UserId = testUser.Id
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
