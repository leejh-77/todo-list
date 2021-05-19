package users

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
	"todo-list/constants"
)

func TestSignup_valid(t *testing.T) {
	c := SignUpCommand{}
	c.EmailAddress = "jonghoon.lee@gmail.com"
	c.Password = "pasworkd@@#"
	c.Username = "Jonghoon Lee"

	token, err := SignUp(c)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, token)
}

func TestSignup_invalid_email(t *testing.T) {
	c := SignUpCommand{}
	c.EmailAddress = "jonghoon.lee"
	c.Password = "pasword@!@!"
	c.Username = "Jonghoon Lee"

	_, err := SignUp(c)
	if err == nil {
		t.Fatal()
	}
	assert.EqualError(t, err, "invalid emailAddress")
}

func TestSignup_invalid_password(t *testing.T) {
	c := SignUpCommand{}
	c.EmailAddress = "jonghoon.lee@gmail.com"
	c.Password = "sss"
	c.Username = "Jonghoon Lee"

	_, err := SignUp(c)
	if err == nil {
		t.Fatal()
	}
	assert.EqualError(t, err, "invalid password")
}

func TestSignUp_validate_jwt(t *testing.T) {
	c := SignUpCommand{}
	c.EmailAddress = "jonghoon.lee@gmail.com"
	c.Password = "password@!$"
	c.Username = "Jonghoon Lee"

	jwtToken, err := SignUp(c)
	if err != nil {
		t.Fatal(err)
	}

	token, err := jwt.Parse(*jwtToken, func(token2 *jwt.Token) (interface{}, error) {
		return []byte(constants.JWTSecret), nil
	})
	if err != nil {
		t.Fatal(err)
	}
	claim := token.Claims.(jwt.MapClaims)
	email := claim["email"]
	assert.Equal(t, c.EmailAddress, email)
}

func TestSignUp_password_encrypt(t *testing.T) {
	email := "jonghoon.lee@gmail.com"
	password := "password@!@"

	c := SignUpCommand{}
	c.EmailAddress = email
	c.Password = password
	c.Username = "Jonghoon Lee"

	_, err := SignUp(c)
	if err != nil {
		t.Fatal(err)
	}

	user, err := findUserByEmailAddress(email)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, user.Password, password)
}