package orm

import (
	"reflect"
	"testing"
)

type TestObject struct {
	StringValue string
	IntValue int
}

func TestReflect(t *testing.T) {
	m := make(map[string]interface{})
	m["StringValue"] = "Jonghoon"
	m["IntValue"] = int64(123)

	obj := TestObject{}

	obj2 := reflect.New(reflect.TypeOf(obj)).Interface()
	val := reflect.ValueOf(obj2).Elem()
	ref := reflect.TypeOf(obj2).Elem()

	t.Log(val.String())

	num := val.NumField()
	for i := 0; i < num; i++ {
		f := val.Field(i)
		rf := ref.Field(i)
		name := rf.Name
		t.Log(name)
		v := m[name]

		k := f.Kind()
		if k == reflect.String {
			f.SetString(v.(string))
		} else if k == reflect.Int {
			f.SetInt(v.(int64))
		}
	}

	t.Log(obj)
}

func TestCreateObject(t *testing.T) {
	obj := TestObject{}
	o := reflect.New(reflect.TypeOf(obj))
	_, ok := o.Elem().Interface().(TestObject)
	t.Log(ok)
}
