package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo-list/test"
)

func init() {
	test.BeforeTest()
}

func TestCreateUser(t *testing.T) {
	email := test.UniqueString("test.user@gmail.com")

	id, err := createUserWithEmail(email)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, -1, id)
}

func TestFindByEmailAddress(t *testing.T) {
	email := test.UniqueString("test.user@gmail.com")

	id, err := createUserWithEmail(email)
	if err != nil {
		t.Fatal(err)
	}
	found, err := FindUserByEmailAddress(email)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, id, found.Id)
}

func TestFindById(t *testing.T) {
	email := test.UniqueString("test.user@gmail.com")

	id, err := createUserWithEmail(email)
	if err != nil {
		t.Fatal(err)
	}
	user, err := FindUserById(id)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, user)
}

func createUserWithEmail(email string) (int64, error) {
	user := new(User)
	user.EmailAddress = email
	user.Password = "passwod@!!"
	user.Username = "Jonghoon Lee"
	user.RegisteredTime = time.Now().Unix()
	return CreateUser(user)
}
