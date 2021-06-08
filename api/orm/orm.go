package orm

import (
	"bytes"
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
	deleteQuery string
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
	t.makeUpdateQuery()
	t.makeDeleteQuery()
}

func (t *Table) makeInsertQuery() {
	params := strings.Repeat("?,", len(t.fields))
	params = params[:len(params) - 1]
	t.insertQuery =
		"INSERT INTO " + t.name + " (" + strings.Join(t.fields, ",") + ") VALUES (" + params + ")"
}

func (t *Table) makeUpdateQuery() {
	buf := bytes.Buffer{}
	buf.WriteString("UPDATE " + t.name + " SET")

	fs := t.fields
	for _, f := range fs {
		buf.WriteString(" " + f + " = ?,")
	}
	buf.Truncate(buf.Len() - 1)
	buf.WriteString(" WHERE id = ?")
	t.updateQuery = buf.String()
}

func (t *Table) makeDeleteQuery() {
	t.deleteQuery = "DELETE FROM " + t.name + " WHERE id = ?"
}

func (t *Table) FindById(i interface{}, id int64) error {
	return t.FindOne(i, "id = ?", id)
}

func (t *Table) FindAll(i interface{}) error {
	return t.Find(i, "")
}

func (t *Table) Find(i interface{}, where string, args... interface{}) error {
	t.ensureType(i)

	q := "SELECT * FROM " + t.name
	if len(where) > 0 {
		q = q + " WHERE " + where
	}
	ret, err := db.Query(q, args...)
	if err != nil {
		return err
	}
	defer ret.Close()
	return t.read(i, ret)
}

func (t *Table) ensureType(i interface{}) {
	k := reflect.TypeOf(i).Kind()
	if k != reflect.Ptr {
		panic("parameter must be pointer")
	}
}

func (t *Table) FindOne(i interface{}, where string, args... interface{}) error {
	err := t.Find(i, where, args...)
	if err != nil {
		return err
	}
	return nil
}

func (t *Table) Insert(i interface{}) (int64, error) {
	params := t.resolveParams(i)
	ret, err := db.Exec(t.insertQuery, params...)
	if err != nil {
		return -1, err
	}
	return ret.LastInsertId()
}

func (t *Table) Update(i interface{}) error {
	id := t.getId(i)
	params := t.resolveParams(i)
	params = append(params, id)

	_, err := db.Exec(t.updateQuery, params...)
	return err
}

func (t *Table) Delete(id int64) error {
	_, err := db.Exec(t.deleteQuery, id)
	return err
}

func (t *Table) getId(i interface{}) int64 {
	v := reflect.ValueOf(i).Elem()
	f := v.FieldByName("Id")
	return f.Int()
}

func (t *Table) resolveParams(i interface{}) []interface{} {
	v := reflect.ValueOf(i).Elem()

	fs := t.fields
	arr := make([]interface{}, 0, len(fs))
	for _, f := range fs {
		f := v.FieldByName(f)
		arr = append(arr, f.Interface())
	}
	return arr
}

func (t *Table) read(i interface{}, row *sql.Rows) error {
	arr := make([]interface{}, 0)
	for row.Next() {
		obj, err := t.readFromRow(row)
		if err != nil {
			return err
		}
		arr = append(arr, obj)
	}

	v := reflect.ValueOf(i).Elem()
	k := v.Kind()

	var refValue reflect.Value
	if len(arr) == 0 {
		refValue = reflect.Zero(v.Type())
	} else if k == reflect.Struct {
		refValue = reflect.ValueOf(arr[0]).Elem()
	} else if k == reflect.Slice {
		refValue = reflect.ValueOf(arr).Elem()
	}
	v.Set(refValue)
	return nil
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

