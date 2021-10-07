package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"log"
	"net/http/httptest"
	"time"
	"todo-list/orm"
)

func ClearTables() {
	_ = orm.Table(TableFolder).DeleteAll()
	_ = orm.Table(TableTodo).DeleteAll()
	_ = orm.Table(TableUser).DeleteAll()
	_ = orm.Table(TableWorkspace).DeleteAll()
	_ = orm.Table(TableWorkspaceMember).DeleteAll()
}

func CreateDummyContext() echo.Context {
	return CreateContext(nil)
}

func CreateAuthorizedContext() echo.Context {
	return CreateContext(TestUser())
}

func CreateContext(user *User) echo.Context {
	e := echo.New()
	req := httptest.NewRequest("GET", "http://localhost", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	if user != nil {
		c := jwt.MapClaims{}
		c["email"] = user.EmailAddress
		c["uid"] = user.Id
		t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
		ctx.Set("user", t)
	}
	return ctx
}

func TestUser() *User {
	email := "todo.test.user@gmail.com"
	return CreateTestUser(email)
}

func CreateTestUser(email string) *User {
	var u User
	_ = UserQuery(orm.Engine).FindByEmailAddress(&u, email)
	if u.Id > 0 {
		return &u
	}
	u = User{
		EmailAddress:   email,
		Password:       "password!@#$",
		Username:       email,
		RegisteredTime: time.Now().Unix(),
	}
	id, err := orm.Table(TableUser).Insert(&u)
	if err != nil {
		log.Fatal(err)
	}
	u.Id = id
	return &u
}

func TestWorkspace() *Workspace {
	name := "test workspace"
	user := TestUser()

	var w Workspace
	err := orm.Table(TableWorkspace).Find(&w,
		"name = ? AND id IN (SELECT workspaceId FROM workspaceMembers WHERE userId = ?)",
		name,
		user.Id)
	if err != nil {
		log.Fatal(err)
	}
	if w.Id != int64(0) {
		return &w
	}

	w = Workspace{
		Name:        name,
		CreatedTime: time.Now().Unix(),
	}
	id, err := orm.Table(TableWorkspace).Insert(&w)
	if err != nil {
		log.Fatal(err)
	}
	w.Id = id
	m := WorkspaceMember{
		Type:        MemberTypeOwner,
		WorkspaceId: w.Id,
		UserId:      user.Id,
	}
	id, err = orm.Table(TableWorkspaceMember).Insert(&m)
	if err != nil {
		log.Fatal(err)
	}
	return &w
}

func TestFolder() *Folder {
	w := TestWorkspace()
	f := &Folder{
		Name: "test.folder",
		WorkspaceId: w.Id,
	}
	id, err := orm.Table(TableFolder).Insert(f)
	if err != nil {
		log.Fatal(err)
	}
	f.Id = id
	return f
}

func TestTodo() *Todo {
	f := TestFolder()
	u := TestUser()
	t := &Todo{
		FolderId:      f.Id,
		UserId:        u.Id,
		Subject:       "test todo",
		Body:          "test todo body",
		Status:        TodoStatusNotStarted,
		CompletedTime: 0,
		Position:      0,
	}
	id, err := orm.Table(TableTodo).Insert(t)
	if err != nil {
		log.Fatal(err)
	}
	t.Id = id
	return t
}
