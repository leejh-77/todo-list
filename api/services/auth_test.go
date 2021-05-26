package services

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-list/models"
	"todo-list/result"
	"todo-list/test"
)

func init() {
	models.BeforeTest()
}

func TestSignup_valid(t *testing.T) {
	ret := signUpTestUser(
		test.UniqueString("jonghoon.lee@gmail.com"),
		"pasworkd@@#",
		"Jonghoon Lee")
	assert.Equal(t, http.StatusCreated, ret.StatusCode)
}

func TestSignup_invalid_email(t *testing.T) {
	ret := signUpTestUser(
		"jonghoon.lee",
		"pasword@!@!",
		"Jonghoon Lee")
	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
}

func TestSignup_invalid_password(t *testing.T) {
	ret := signUpTestUser(
		"jonghoon.lee@gmail.com",
		"sss",
		"Jonghoon Lee")
	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
}

func TestSignUp_password_encrypt(t *testing.T) {
	email := test.UniqueString("jonghoon.lee@gmail.com")
	password := "password@!@"

	ret := signUpTestUser(
		email,
		password,
		"Jonghoon Lee")

	assert.Equal(t, http.StatusCreated, ret.StatusCode)

	user, err := models.FindUserByEmailAddress(email)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, user.Password, password)
}

func TestLogin_valid(t *testing.T) {
	email, password := test.UniqueString("test.user@gmail.com"), "password"
	ret := signUpTestUser(email, password, "Test User")

	assert.Equal(t, http.StatusCreated, ret.StatusCode)

	c := LogInCommand{
		EmailAddress: email,
		Password: password,
	}
	ctx := dummyContext()
	ctx.Set("command", c)
	ret = LogIn(ctx)

	assert.Equal(t, http.StatusOK, ret.StatusCode)

	token := ctx.Response().Header().Get("Set-Cookie")
	assert.NotNil(t, token)
}

func TestLogin_should_fail(t *testing.T) {
	email, password := "test.user@gmail.com", "password"
	ret := signUpTestUser(email, password, "Test User")

	assert.Equal(t, http.StatusCreated, ret.StatusCode)

	c := LogInCommand{
		EmailAddress: email,
		Password: password + "1",
	}

	ctx := dummyContext()
	ctx.Set("command", c)
	ret = LogIn(ctx)
	assert.Equal(t, http.StatusBadRequest, ret.StatusCode)
}

func dummyContext() echo.Context {
	req := httptest.NewRequest("GET", "http://localhost", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	return e.NewContext(req, rec)
}

func signUpTestUser(email string, password string, username string) *result.ApiResult {
	c := SignUpCommand{}
	c.EmailAddress = email
	c.Password = password
	c.Username = username

	ctx := dummyContext()
	ctx.Set("command", c)
	return SignUp(ctx)
}
