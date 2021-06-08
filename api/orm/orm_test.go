package orm

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Book struct {
	Id int64
	Subject string
	Author string
	PublishedTime time.Time
}

var table = NewTable("books", Book{})

func TestCreate(t *testing.T) {
	book := bookMock()
	id, err := table.Insert(book)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, id, int64(-1))
}

func TestUpdate(t *testing.T) {
	book := bookMock()
	id, err := table.Insert(book)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, id, int64(-1))

	err = table.FindById(book, id)
	if err != nil {
		t.Fatal(err)
	}
	book.Author = "Another author"
	err = table.Update(book)
	if err != nil {
		t.Fatal(err)
	}
	err = table.FindById(book, id)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, book.Author, "Another author")
}

func TestDelete(t *testing.T) {
	book := bookMock()
	id, err := table.Insert(book)
	if err != nil {
		t.Fatal(err)
	}
	err = table.Delete(id)
	if err != nil {
		t.Fatal(err)
	}
	err = table.FindById(book, id)
	if err != nil {
		t.Fatal()
	}
	assert.Equal(t, int64(0), book.Id)
}

func TestFind(t *testing.T) {
	arr := make([]*Book, 0)
	err := table.FindAll(arr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(arr)
}

func bookMock() *Book {
	return &Book{
		Subject:   "fun programming",
		Author:       "programmer",
		PublishedTime: time.Now(),
	}
}

