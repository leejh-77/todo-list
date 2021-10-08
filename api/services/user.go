package services

import (
	"todo-list/models"
	"todo-list/orm"
	"todo-list/result"
)

type GetUserResponse struct {
	Id int64 `json:"id"`
	EmailAddress string `json:"emailAddress"`
	Username string `json:"username"`
}

func responseFromUser(u models.User) *GetUserResponse {
	return &GetUserResponse{
		Id: u.Id,
		EmailAddress: u.EmailAddress,
		Username: u.Username,
	}
}

func GetUser(uid int64) *result.ApiResult {
	var u models.User
	err := orm.Table(models.TableUser).FindById(&u, uid)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(responseFromUser(u))
}
