package controllers

import (
	"github.com/labstack/echo/v4"
	"todo-list/result"
	"todo-list/services"
)

type AuthController struct {

}

func (c AuthController) Init(g *echo.Group) {
	// auth
	g.POST("/signup", signUp, withCommand(services.SignUpCommand{}))
	g.POST("/login", logIn, withCommand(services.LogInCommand{}))
}

func signUp(ctx echo.Context) error {
	return send(ctx, services.SignUp(ctx))
}

func logIn(ctx echo.Context) error {
	return send(ctx, services.LogIn(ctx))
}

func send(ctx echo.Context, r *result.ApiResult) error {
	err := r.Error
	if err != nil && err.Error != nil {
		return err.Error
	}
	return ctx.JSON(r.StatusCode, r.Result)
}

func withCommand(i interface{}) echo.MiddlewareFunc {
	return func(fn echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			err := ctx.Bind(&i)
			if err != nil {
				return err
			}
			ctx.Set("command", i)
			return fn(ctx)
		}
	}
}
