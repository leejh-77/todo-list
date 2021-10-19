package controllers

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"todo-list/services"
)

type UserController struct {

}

func (c UserController) Init(g *echo.Group)  {
	g.GET("/:id", getUser)
	g.PUT("", updateUser)
}

func getUser(ctx echo.Context) error {
	param := ctx.Param("id")
	var uid int64
	uid = userIdFromContext(ctx)

	if param == "me" {
		return services.GetMe(uid).Send(ctx)
	}

	v, err := strconv.Atoi(param)
	if err != nil {
		return err
	}
	uid = int64(v)
	return services.GetUser(uid).Send(ctx)
}

func updateUser(ctx echo.Context) error {
	uid := userIdFromContext(ctx)
	var c services.UpdateUserCommand
	err := ctx.Bind(&c)
	if err != nil {
		return err
	}
	return services.UpdateUser(uid, c).Send(ctx)
}
