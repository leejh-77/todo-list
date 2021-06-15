package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"todo-list/models"
	"todo-list/orm"
	"todo-list/result"
	"todo-list/test"
)

func TestSignup(t *testing.T) {
	ret := signUpTestUser(
		test.UniqueString("jonghoon.lee@gmail.com"),
		"pasworkd@@#",
		"Jonghoon Lee")
	assert.Equal(t, http.StatusCreated, ret.StatusCode)
}

func TestSignup_invalidEmail_shouldFail(t *testing.T) {
	ret := signUpTestUser(
		"jonghoon.lee",
		"pasword@!@!",
		"Jonghoon Lee")
	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
}

func TestSignup_invalidPassword_shouldFail(t *testing.T) {
	ret := signUpTestUser(
		"jonghoon.lee@gmail.com",
		"sss",
		"Jonghoon Lee")
	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
}

func TestSignUp_passwordEncrypt(t *testing.T) {
	email := test.UniqueString("jonghoon.lee@gmail.com")
	password := "password@!@"

	ret := signUpTestUser(
		email,
		password,
		"Jonghoon Lee")

	assert.Equal(t, http.StatusCreated, ret.StatusCode)

	var user models.User
	err := orm.Table(models.TableUser).FindOne(&user, "emailAddress = ?", email)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, user.Password, password)
}

func TestLogin(t *testing.T) {
	email, password := test.UniqueString("test.user@gmail.com"), "password"
	ret := signUpTestUser(email, password, "Test User")

	assert.Equal(t, http.StatusCreated, ret.StatusCode)

	c := LogInCommand{
		EmailAddress: email,
		Password: password,
	}
	ctx := createDummyContext()
	ret = LogIn(ctx, c)

	assert.Equal(t, http.StatusOK, ret.StatusCode)

	token := ctx.Response().Header().Get("Set-Cookie")
	assert.NotNil(t, token)
}

func TestLogin_invalidPassword_shouldFail(t *testing.T) {
	email, password := "test.user@gmail.com", "password"
	ret := signUpTestUser(email, password, "Test User")

	assert.Equal(t, http.StatusCreated, ret.StatusCode)

	c := LogInCommand{
		EmailAddress: email,
		Password: password + "1",
	}
	ctx := createDummyContext()
	ret = LogIn(ctx, c)
	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
}


func signUpTestUser(email string, password string, username string) *result.ApiResult {
	c := SignUpCommand{}
	c.EmailAddress = email
	c.Password = password
	c.Username = username
	return SignUp(c)
}
