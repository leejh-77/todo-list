package orm

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var config = DatabaseConfig{
	Driver:   "mysql",
	Host:     "127.0.0.1:3306",
	User:     "root",
	Password: "1234",
	Name:     "orm",
}

type Book struct {
	Id int64
	Subject string
	Author string
	PublishedTime int64
}

var tBook = "books"

func init() {
	Init(config)
	Register(tBook, Book{})
}

func TestCreate(t *testing.T) {
	book := bookMock()
	id, err := Table(tBook).Insert(book)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, id, int64(-1))
}

func TestUpdate(t *testing.T) {
	book := bookMock()
	id, err := Table(tBook).Insert(book)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, id, int64(-1))

	err = Table(tBook).FindById(book, id)
	if err != nil {
		t.Fatal(err)
	}
	book.Author = "Another author"
	err = Table(tBook).Update(book)
	if err != nil {
		t.Fatal(err)
	}
	err = Table(tBook).FindById(book, id)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, book.Author, "Another author")
}

func TestDelete(t *testing.T) {
	book := bookMock()
	id, err := Table(tBook).Insert(book)
	if err != nil {
		t.Fatal(err)
	}
	err = Table(tBook).Delete(id)
	if err != nil {
		t.Fatal(err)
	}
	err = Table(tBook).FindById(book, id)
	if err != nil {
		t.Fatal()
	}
	assert.Equal(t, int64(0), book.Id)
}

func TestFind(t *testing.T) {
	var books []Book
	err := Table(tBook).FindAll(&books)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(books)
}

func bookMock() *Book {
	return &Book{
		Subject:   "fun programming",
		Author:       "programmer",
		PublishedTime: time.Now().Unix(),
	}
}

