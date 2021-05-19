package auth

import (
	"github.com/labstack/echo/v4"
	"todo-list/result"
)

type SignUpCommand struct {
	EmailAddress string `json:"emailAddress"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type LogInCommand struct {
	EmailAddress string `json:"emailAddress"`
	Password string `json:"password"`
}

func Init(g *echo.Group) {
	g.POST("/signup", doSignUp, withCommand(SignUpCommand{}))
	g.POST("/login", doLogin, withCommand(LogInCommand{}))
}

func doSignUp(ctx echo.Context) error {
	return send(ctx, signUp(ctx.Get("command").(SignUpCommand)))
}

func doLogin(ctx echo.Context) error {
	return send(ctx, login(ctx, ctx.Get("command").(LogInCommand)))
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
