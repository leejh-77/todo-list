package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo-list/orm"
	"todo-list/test"
)

func TestCreateUser(t *testing.T) {
	email := test.UniqueString("test.user@gmail.com")

	id, err := createUser(email)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, -1, id)
}

func TestFindUserByEmailAddress(t *testing.T) {
	email := test.UniqueString("test.user@gmail.com")

	id, err := createUser(email)
	if err != nil {
		t.Fatal(err)
	}
	u := &User{}
	err = UserQuery(orm.Engine).FindByEmailAddress(u,  email)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, id, u.Id)
}

func TestFindUserById(t *testing.T) {
	email := test.UniqueString("test.user@gmail.com")

	id, err := createUser(email)
	if err != nil {
		t.Fatal(err)
	}

	u := &User{}
	err = orm.Table(TableUser).FindById(u, id)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, u)
}

func createUser(email string) (int64, error) {
	user := new(User)
	user.EmailAddress = email
	user.Password = "passwod@!!"
	user.Username = "Jonghoon Lee"
	user.RegisteredTime = time.Now().Unix()
	return orm.Table(TableUser).Insert(user)
}

