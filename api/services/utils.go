package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)


func userIdFromContext(ctx echo.Context) int64 {
	t := ctx.Get("user").(*jwt.Token)
	claims := t.Claims.(jwt.MapClaims)
	uid := claims["uid"].(int64)
	return uid
}