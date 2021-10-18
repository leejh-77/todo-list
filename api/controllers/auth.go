package controllers

import (
	"github.com/labstack/echo/v4"
	"todo-list/services"
)

type AuthController struct {

}

func (c AuthController) Init(g *echo.Group) {
	// auth
	g.POST("/signup", signUp)
	g.POST("/login", logIn)
	g.POST("/logout", logOut)
}

func signUp(ctx echo.Context) error {
	var c services.SignUpCommand
	err := ctx.Bind(&c)
	if err != nil {
		return err
	}
	return services.SignUp(c).Send(ctx)
}

func logIn(ctx echo.Context) error {
	var c services.LogInCommand
	err := ctx.Bind(&c)
	if err != nil {
		return err
	}
	return services.LogIn(ctx, c).Send(ctx)
}

func logOut(ctx echo.Context) error {
	return services.LogOut(ctx).Send(ctx)
}
