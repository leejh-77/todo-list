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
}

func getUser(ctx echo.Context) error {
	param := ctx.Param("id")
	var uid int64
	if param == "me" {
		uid = userIdFromContext(ctx)
	} else {
		v, err := strconv.Atoi(param)
		if err != nil {
			return err
		}
		uid = int64(v)
	}
	return services.GetUser(uid).Send(ctx)
}
