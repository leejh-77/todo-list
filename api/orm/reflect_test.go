package orm

import (
	"log"
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

func TestReflectArray(t *testing.T) {
	var i int64
	fillData(&i)

	var arr []int64
	fillData(&arr)
}

func fillData(i interface{}) {
	data := make([]int64, 0)
	for i := 5; i > 0; i-- {
		data = append(data, int64(i))
	}

	v := reflect.ValueOf(i).Elem()
	k := v.Kind()
	if k == reflect.Int {
		v.SetInt(data[0])
	} else if k == reflect.Slice {
		v.Set(reflect.ValueOf(data))
	}
	log.Println(k)
}
