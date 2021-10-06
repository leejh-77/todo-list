package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func userIdFromContext(ctx echo.Context) int64 {
	t := ctx.Get("user").(*jwt.Token)
	claims := t.Claims.(jwt.MapClaims)
	uid := claims["uid"].(float64) // int64로 값을 넣어줬음에도 float64 로 돌려준다. json 파싱 라이브러리 때문인듯
	return int64(uid)
}