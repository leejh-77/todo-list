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

type DatabaseConfig struct {
	Driver string
	Host string
	User string
	Password string
	Name string
}

func Init(c DatabaseConfig) {
	i, err := sql.Open(c.Driver, c.User + ":" + c.Password + "@tcp(" + c.Host + ")/" + c.Name)
	if err != nil {
		log.Fatal(err)
	}
	db = i
}

type Table struct {
	name string
	entityType reflect.Type
	fields []reflect.StructField

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
	fs := make([]reflect.StructField, 0, cnt)
	for i := 0; i < cnt; i++ {
		f := e.Field(i)
		n := f.Name
		if n == "Id" {
			continue
		}
		fs = append(fs, f)
	}
	t.fields = fs

	t.createTable()

	t.makeInsertQuery()
	t.makeUpdateQuery()
	t.makeDeleteQuery()
}

func (t *Table) createTable() {
	buf := bytes.Buffer{}
	buf.WriteString("CREATE TABLE IF NOT EXISTS `" + t.name + "` (")
	buf.WriteString("`id` INT PRIMARY KEY AUTO_INCREMENT, ")

	for _, f := range t.fields {
		str := t.typeString(f.Type)
		buf.WriteString("`" + f.Name + "` " + str + ", ")
	}
	buf.Truncate(buf.Len() - 2)
	buf.WriteString(")")

	query := buf.String()
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func (t *Table) typeString(p reflect.Type) string {
	k := p.Kind()
	if k == reflect.String {
		return "TEXT"
	}
	if k == reflect.Int64 {
		return "INT"
	}
	str := p.String()
	if str == "time.Time" {
		return "TIMESTAMP"
	}
	return "BLOB"
}

func (t *Table) makeInsertQuery() {
	buf := bytes.Buffer{}
	buf.WriteString("INSERT INTO " + t.name + " (")

	fs := t.fields
	for _, f := range fs {
		buf.WriteString(f.Name + ",")
	}
	buf.Truncate(buf.Len() - 1)
	buf.WriteString(") VALUES (")

	params := strings.Repeat("?,", len(t.fields))
	params = params[:len(params) - 1]

	buf.WriteString(params)
	buf.WriteString(")")
	t.insertQuery = buf.String()
}

func (t *Table) makeUpdateQuery() {
	buf := bytes.Buffer{}
	buf.WriteString("UPDATE " + t.name + " SET")

	fs := t.fields
	for _, f := range fs {
		buf.WriteString(" " + f.Name + " = ?,")
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
		f := v.FieldByName(f.Name)
		arr = append(arr, f.Interface())
	}
	return arr
}

func (t *Table) read(i interface{}, row *sql.Rows) error {
	arr := reflect.MakeSlice(reflect.SliceOf(t.entityType), 0, 20)
	for row.Next() {
		obj, err := t.readFromRow(row)
		if err != nil {
			return err
		}
		arr = reflect.Append(arr, reflect.ValueOf(obj).Elem())
	}

	v := reflect.ValueOf(i).Elem()
	k := v.Kind()

	var refValue reflect.Value
	if arr.Len() == 0 {
		refValue = reflect.Zero(v.Type())
	} else if k == reflect.Struct {
		refValue = arr.Index(0)
	} else if k == reflect.Slice {
		refValue = arr
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

