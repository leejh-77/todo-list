package services

import (
	"todo-list/models"
	"todo-list/orm"
	"todo-list/result"
	"todo-list/utils"
)

type GetUserResponse struct {
	Id int64 `json:"id"`
	EmailAddress string `json:"emailAddress"`
	Username string            `json:"username"`
	Image    *utils.Base64Image `json:"image"`
}

type UpdateUserCommand struct {
	Image    *utils.Base64Image `json:"image"`
	Username string            `json:"username"`
}

func GetUser(uid int64) *result.ApiResult {
	var u models.User
	err := orm.Table(models.TableUser).FindById(&u, uid)
	if err != nil {
		return result.ServerError(err)
	}
	return userResponse(u)
}

func UpdateUser(uid int64, c UpdateUserCommand) *result.ApiResult {
	if c.Image != nil {
		err := utils.WriteImage(uid, c.Image)
		if err != nil {
			return result.ServerError(err)
		}
	}

	var u models.User
	err := orm.Table(models.TableUser).FindById(&u, uid)
	if err != nil {
		return result.ServerError(err)
	}
	u.Username = c.Username
	err = orm.Table(models.TableUser).Update(&u)
	if err != nil {
		return result.ServerError(err)
	}
	return userResponse(u)
}

func userResponse(u models.User) *result.ApiResult {
	res := &GetUserResponse{
		Id: u.Id,
		EmailAddress: u.EmailAddress,
		Username: u.Username,
	}
	image, err := utils.ReadImage(u.Id)
	if err != nil {
		return result.ServerError(err)
	}
	res.Image = image
	return result.Success(res)
}

