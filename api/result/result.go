package result

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type ApiResult struct {
	Result interface{} `json:"result"`
	StatusCode int     `json:"statusCode"`
	Error *ApiError    `json:"error"`
}

type ApiError struct {
	Message string `json:"message"`
	Error error `json:"-"`
}

func Success(result interface{}) *ApiResult {
	return &ApiResult{
		Result: result,
		StatusCode: http.StatusOK,
	}
}

func Created() *ApiResult {
	return &ApiResult{
		Result: "",
		StatusCode: http.StatusCreated,
	}
}

func Error(code int, message string) *ApiResult {
	return &ApiResult{
		Result: message,
		StatusCode: code,
	}
}

func BadRequest(message string) *ApiResult {
	return &ApiResult{
		Result: message,
		StatusCode: http.StatusBadRequest,
	}
}

func Unauthorized(message string) *ApiResult {
	return &ApiResult{
		Result: message,
		StatusCode: http.StatusUnauthorized,
	}
}

func ServerError(err error) *ApiResult {
	return &ApiResult{
		StatusCode: http.StatusInternalServerError,
		Error: &ApiError{
			Message: "Internal server error",
			Error: err,
		},
	}
}

func (ret *ApiResult) Send(ctx echo.Context) error {
	err := ret.Error
	if err != nil && err.Error != nil {
		return err.Error
	}
	log.Println(ret.Result)
	header := ctx.Response().Header()
	header.Set("Cache-Control", "no-cache")
	header.Set("Pragma", "no-cache")
	return ctx.JSON(ret.StatusCode, ret.Result)
}

