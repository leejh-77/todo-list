package model

type User struct {
	Id int64
	EmailAddress string
	Password string
	Username string
	RegisteredTime int64
}

const (
	TodoStatusNotStarted = 0
	TodoStatusInProgress = 1
	TodoStatusCompleted = 2
)

type Todo struct {
	Id int64
	UserId int64
	Subject string
	Body string
	Status int
	CompletedTime int64
}
