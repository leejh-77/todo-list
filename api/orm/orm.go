package orm

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
	"strings"
)

var db *sql.DB

func init() {
	d, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/todo")
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

type Table struct {
	name string
	entityType reflect.Type
	fields []string

	insertQuery string
	updateQuery string
}

func NewTable(name string, entity interface{}) *Table {
	table := &Table{}
	table.name = name
	table.entityType = reflect.TypeOf(entity)
	table.init()
	return table
}

func (t *Table) init() {
	e := t.entityType
	cnt := e.NumField()
	fs := make([]string, 0, cnt)
	for i := 0; i < cnt; i++ {
		n := e.Field(i).Name
		if n == "Id" {
			continue
		}
		fs = append(fs, n)
	}
	t.fields = fs

	t.makeInsertQuery()
}

func (t *Table) makeInsertQuery() {
	params := strings.Repeat("?,", len(t.fields))
	params = params[:len(params) - 1]
	t.insertQuery =
		"INSERT INTO " + t.name + " (" + strings.Join(t.fields, ",") + ") VALUES (" + params + ")"
}

func (t *Table) FindById(id int64) (interface{}, error) {
	return t.FindOne("id = ?", id)
}

func (t *Table) FindAll() ([]interface{}, error) {
	return t.Find("")
}

func (t *Table) Find(where string, args... interface{}) ([]interface{}, error) {
	q := "SELECT * FROM " + t.name
	if len(where) > 0 {
		q = q + " WHERE " + where
	}
	ret, err := db.Query(q, args...)
	if err != nil {
		return nil, err
	}
	defer ret.Close()
	return t.read(ret)
}

func (t *Table) FindOne(where string, args... interface{}) (interface{}, error) {
	arr, err := t.Find(where, args)
	if err != nil {
		return nil, err
	}
	if len(arr) == 0 {
		return nil, nil
	}
	return arr[0], nil
}

func (t *Table) Save(i interface{}) (int64, error) {
	v := reflect.ValueOf(i)

	cnt := v.NumField()
	arr := make([]interface{}, 0, cnt)
	for i := 0; i < cnt; i++ {
		if i == 0 { // 땜빵
			continue
		}
		f := v.Field(i)
		arr = append(arr, f.Interface())
	}
	ret, err := db.Exec(t.insertQuery, arr...)
	if err != nil {
		return -1, err
	}
	return ret.LastInsertId()
}

func (t *Table) read(row *sql.Rows) ([]interface{}, error) {
	arr := make([]interface{}, 0)
	for row.Next() {
		obj, err := t.readFromRow(row)
		if err != nil {
			return nil, err
		}
		arr = append(arr, obj)
	}
	return arr, nil
}

func (t *Table) readFromRow(row *sql.Rows) (interface{}, error){
	obj := reflect.New(t.entityType).Interface()
	v := reflect.ValueOf(obj).Elem()

	cnt := v.NumField()
	fs := make([]interface{}, 0, cnt)
	for i := 0; i < cnt; i++ {
		f := v.Field(i).Addr().Interface()
		fs = append(fs, f)
	}
	err := row.Scan(fs...)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

