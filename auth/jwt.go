package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"todo-list/constants"
)

func CreateJwt(email string) (*string, error) {
	c := jwt.MapClaims{}
	c["exp"] = time.Now().Add(time.Hour).Unix()
	c["email"] = email

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	token, err := t.SignedString([]byte(constants.JWTSecret))
	if err != nil {
		return nil, err
	}
	return &token, nil
}
