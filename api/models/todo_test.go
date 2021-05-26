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
	id, err := CreateUser(u)
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
	todos, err := FindAll()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, len(todos))
}

func TestCreate(t *testing.T) {
	todo := createTestTodo()
	assert.NotNil(t, todo)
}

func TestUpdate(t *testing.T) {
	todo := createTestTodo()
	todo.Subject = "Updated subject"

	_, err := SaveTodo(todo)
	todo, err = FindById(todo.Id)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Updated subject", todo.Subject)
}

func TestDelete(t *testing.T) {
	todo := createTestTodo()
	err := DeleteTodo(todo.Id)
	if err != nil {
		t.Fatal(err)
	}
	todo, err = FindById(todo.Id)
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
	id, err := SaveTodo(todo)
	if err != nil {
		return nil
	}
	todo, err = FindById(id)
	if err != nil {
		return nil
	}
	return todo
}
