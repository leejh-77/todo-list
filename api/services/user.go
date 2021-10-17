package services

import (
	base642 "encoding/base64"
	"io/ioutil"
	"strconv"
	"todo-list/models"
	"todo-list/orm"
	"todo-list/result"
)

type GetUserResponse struct {
	Id int64 `json:"id"`
	EmailAddress string `json:"emailAddress"`
	Username string `json:"username"`
}

type UpdateUserCommand struct {
	ImageData string `json:"imageData"`
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

func UpdateUser(uid int64, c UpdateUserCommand) *result.ApiResult {
	base64 := c.ImageData
	decoded, err := base642.StdEncoding.DecodeString(base64)
	if err != nil {
		return result.ServerError(err)
	}
	err = ioutil.WriteFile("../profile/" + strconv.FormatInt(uid, 10), decoded, 0644)
	if err != nil {
		return result.ServerError(err)
	}

	var u models.User
	err = orm.Table(models.TableUser).FindById(&u, uid)
	if err != nil {
		return result.ServerError(err)
	}
	u.Username = c.Username
	err = orm.Table(models.TableUser).Update(u)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(responseFromUser(u))
}
