package services

import (
	"errors"
	"todo-list/models"
	"todo-list/orm"
	"todo-list/result"
)

type CreateTodoCommand struct {
	FolderId int64
	UserId int64
	Subject string
	Body string
	Status int
	CompletedTime int64
}

func CreateTodo(c CreateTodoCommand) *result.ApiResult {
	err := validateCreateCommand(c)
	if err != nil {
		return result.BadRequest(err.Error())
	}

	todo := &models.Todo{
		FolderId:      c.FolderId,
		UserId:        c.UserId,
		Subject:       c.Subject,
		Body:          c.Body,
		Status:        c.Status,
		CompletedTime: c.CompletedTime,
		Position:      0,
	}
	_, err = orm.Table(models.TableTodo).Insert(todo)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Created()
}

func validateCreateCommand(c CreateTodoCommand) error {
	if c.FolderId == int64(0) {
		return errors.New("folder id must not be empty")
	}
	if c.UserId == int64(0) {
		return errors.New("user id must not be empty")
	}
	if len(c.Subject) == 0 {
		return errors.New("subject must not be empty")
	}
	return nil
}