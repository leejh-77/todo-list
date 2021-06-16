package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"log"
	"net/http/httptest"
	"time"
	"todo-list/models"
	"todo-list/orm"
)

func clearTables() {
	_ = orm.Table(models.TableFolder).DeleteAll()
	_ = orm.Table(models.TableTodo).DeleteAll()
	_ = orm.Table(models.TableUser).DeleteAll()
	_ = orm.Table(models.TableWorkspace).DeleteAll()
	_ = orm.Table(models.TableWorkspaceMember).DeleteAll()
}

func createDummyContext() echo.Context {
	return createContext(nil)
}

func createAuthorizedContext() echo.Context {
	return createContext(TestUser())
}

func createContext(user *models.User) echo.Context {
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

func TestUser() *models.User {
	email := "todo.test.user@gmail.com"
	return createTestUser(email)
}

func createTestUser(email string) *models.User {
	var u models.User
	_ = models.UserQuery(orm.Engine).FindByEmailAddress(&u, email)
	if u.Id > 0 {
		return &u
	}
	u = models.User{
		EmailAddress:   email,
		Password:       "password!@#$",
		Username:       email,
		RegisteredTime: time.Now().Unix(),
	}
	id, err := orm.Table(models.TableUser).Insert(&u)
	if err != nil {
		log.Fatal(err)
	}
	u.Id = id
	return &u
}

func TestWorkspace() *models.Workspace {
	name := "test workspace"
	user := TestUser()

	var w models.Workspace
	err := orm.Table(models.TableWorkspace).Find(&w,
		"name = ? AND id IN (SELECT workspaceId FROM workspaceMembers WHERE userId = ?)",
		name,
		user.Id)
	if err != nil {
		log.Fatal(err)
	}
	if w.Id != int64(0) {
		return &w
	}

	w = models.Workspace{
		Name:        name,
		CreatedTime: time.Now().Unix(),
	}
	id, err := orm.Table(models.TableWorkspace).Insert(&w)
	if err != nil {
		log.Fatal(err)
	}
	w.Id = id
	m := models.WorkspaceMember{
		Type:        models.MemberTypeOwner,
		WorkspaceId: w.Id,
		UserId:      user.Id,
	}
	id, err = orm.Table(models.TableWorkspaceMember).Insert(&m)
	if err != nil {
		log.Fatal(err)
	}
	return &w
}

func TestFolder() *models.Folder {
	w := TestWorkspace()
	f := &models.Folder{
		Name: "test.folder",
		WorkspaceId: w.Id,
	}
	id, err := orm.Table(models.TableFolder).Insert(f)
	if err != nil {
		log.Fatal(err)
	}
	f.Id = id
	return f
}

func TestTodo() *models.Todo {
	f := TestFolder()
	u := TestUser()
	t := &models.Todo{
		FolderId:      f.Id,
		UserId:        u.Id,
		Subject:       "test todo",
		Body:          "test todo body",
		Status:        models.TodoStatusNotStarted,
		CompletedTime: time.Now().Unix(),
		Position:      0,
	}
	id, err := orm.Table(models.TableTodo).Insert(t)
	if err != nil {
		log.Fatal(err)
	}
	t.Id = id
	return t
}
