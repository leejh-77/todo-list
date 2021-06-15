package orm

import (
	"bytes"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strings"
)

type TableInfo struct {
	name string
	entityType reflect.Type
	fields []reflect.StructField

	insertQuery string
	updateQuery string
	deleteQuery string
}

type ORMTable struct {
	engine Session
	info   *TableInfo
}

func (t *TableInfo) init() {
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

	t.makeInsertQuery()
	t.makeUpdateQuery()
	t.makeDeleteQuery()
}

func (t *TableInfo) makeInsertQuery() {
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

func (t *TableInfo) makeUpdateQuery() {
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

func (t *TableInfo) makeDeleteQuery() {
	t.deleteQuery = "DELETE FROM " + t.name + " WHERE id = ?"
}

func (t *ORMTable) FindById(i interface{}, id int64) error {
	return t.FindOne(i, "id = ?", id)
}

func (t *ORMTable) FindAll(i interface{}) error {
	return t.Find(i, "")
}

func (t *ORMTable) Find(i interface{}, where string, args... interface{}) error {
	t.ensureType(i)

	q := "SELECT * FROM " + t.info.name
	if len(where) > 0 {
		q = q + " WHERE " + where
	}
	ret, err := t.engine.Query(q, args...)
	if err != nil {
		return err
	}
	defer ret.Close()
	return t.read(i, ret)
}

func (t *ORMTable) ensureType(i interface{}) {
	k := reflect.TypeOf(i).Kind()
	if k != reflect.Ptr {
		panic("parameter must be pointer")
	}
}

func (t *ORMTable) FindOne(i interface{}, where string, args... interface{}) error {
	err := t.Find(i, where, args...)
	if err != nil {
		return err
	}
	return nil
}

func (t *ORMTable) Insert(i interface{}) (int64, error) {
	params := t.resolveParams(i)
	ret, err := t.engine.Exec(t.info.insertQuery, params...)
	if err != nil {
		return -1, err
	}
	return ret.LastInsertId()
}

func (t *ORMTable) Update(i interface{}) error {
	id := t.getId(i)
	params := t.resolveParams(i)
	params = append(params, id)

	_, err := t.engine.Exec(t.info.updateQuery, params...)
	return err
}

func (t *ORMTable) Delete(where string, i... interface{}) error {
	q := "DELETE FROM " + t.info.name
	if len(where) > 0 {
		q = q + " WHERE " + where
	}
	_, err := t.engine.Exec(q, i...)
	return err
}

func (t *ORMTable) DeleteById(id int64) error {
	_, err := t.engine.Exec(t.info.deleteQuery, id)
	return err
}

func (t *ORMTable) DeleteAll() error {
	_, err := t.engine.Exec("TRUNCATE TABLE " + t.info.name)
	return err
}

func (t *ORMTable) getId(i interface{}) int64 {
	v := reflect.ValueOf(i).Elem()
	f := v.FieldByName("Id")
	return f.Int()
}

func (t *ORMTable) resolveParams(i interface{}) []interface{} {
	v := reflect.ValueOf(i).Elem()

	fs := t.info.fields
	arr := make([]interface{}, 0, len(fs))
	for _, f := range fs {
		f := v.FieldByName(f.Name)
		arr = append(arr, f.Interface())
	}
	return arr
}

func (t *ORMTable) read(i interface{}, row *sql.Rows) error {
	arr := reflect.MakeSlice(reflect.SliceOf(t.info.entityType), 0, 20)
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

func (t *ORMTable) readFromRow(row *sql.Rows) (interface{}, error){
	obj := reflect.New(t.info.entityType).Interface()
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

