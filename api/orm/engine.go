package orm

import (
	"bytes"
	"database/sql"
	"log"
	"reflect"
)

type DatabaseConfig struct {
	Driver string
	Host string
	User string
	Password string
	Name string
}

var tableInfos = make(map[string]*TableInfo)
var Engine *DefaultEngine

type Session interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Table(name string) *ORMTable
}

type DefaultEngine struct {
	db *sql.DB
}

type TransactionalEngine struct {
	db *sql.Tx
}

func (e *TransactionalEngine) Commit() error {
	return e.db.Commit()
}

func (e *TransactionalEngine) Rollback() error {
	return e.db.Rollback()
}

func Init(c DatabaseConfig) {
	i, err := sql.Open(c.Driver, c.User + ":" + c.Password + "@tcp(" + c.Host + ")/" + c.Name)
	if err != nil {
		log.Fatal(err)
	}
	Engine = &DefaultEngine{
		db: i,
	}
}

func (e *DefaultEngine) Exec(query string, args ...interface{}) (sql.Result, error) {
	return e.db.Exec(query, args...)
}

func (e *DefaultEngine) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return e.db.Query(query, args...)
}

func (e *TransactionalEngine) Exec(query string, args ...interface{}) (sql.Result, error) {
	return e.db.Exec(query, args...)
}

func (e *TransactionalEngine) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return e.db.Query(query, args...)
}

func (e *DefaultEngine) Table(name string) *ORMTable {
	return &ORMTable{
		engine: e,
		info:   tableInfos[name],
	}
}

func (e *TransactionalEngine) Table(name string) *ORMTable {
	return &ORMTable{
		engine: e,
		info: tableInfos[name],
	}
}

func Table(name string) *ORMTable {
	return Engine.Table(name)
}

func Register(name string, entity interface{}) {
	info := &TableInfo{
		name:        name,
		entityType:  reflect.TypeOf(entity),
	}
	info.init()
	createTable(info)

	tableInfos[info.name] = info
}

func createTable(info *TableInfo) {
	buf := bytes.Buffer{}
	buf.WriteString("CREATE TABLE IF NOT EXISTS `" + info.name + "` (")
	buf.WriteString("`id` INT PRIMARY KEY AUTO_INCREMENT, ")

	for _, f := range info.fields {
		str := typeString(f.Type)
		buf.WriteString("`" + f.Name + "` " + str + ", ")
	}
	buf.Truncate(buf.Len() - 2)
	buf.WriteString(")")

	query := buf.String()
	_, err := Engine.Exec(query)
	if err != nil {
		panic(err)
	}
}

func typeString(p reflect.Type) string {
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

type TransactionFunc func(Session) error

func InTransaction(fn TransactionFunc) error {
	tx, err := Engine.db.Begin()
	if err != nil {
		return err
	}
	e := &TransactionalEngine{
		db: tx,
	}
	err = fn(e)
	if err == nil {
		err = tx.Commit()
	} else {
		_ = tx.Rollback()
	}
	return err
}

func BeginTr() (*TransactionalEngine, error) {
	tx, err := Engine.db.Begin()
	if err != nil {
		return nil, err
	}
	e := &TransactionalEngine{
		db: tx,
	}
	return e, nil
}