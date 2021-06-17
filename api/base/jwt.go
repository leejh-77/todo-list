package base

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateJwt(id int64, email string) (*string, error) {
	c := jwt.MapClaims{}
	c["exp"] = time.Now().Add(time.Hour).Unix()
	c["email"] = email
	c["uid"] = id

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	token, err := t.SignedString([]byte(JWTSecret))
	if err != nil {
		return nil, err
	}
	return &token, nil
}
