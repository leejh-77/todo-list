package optional

import "encoding/json"

type String struct {
	Set bool
	Value string
}

type Int struct {
	Set bool
	Value int
}

type Int64 struct {
	Set bool
	Value int64
}

func NewString(v string) String {
	return String{
		Set:   true,
		Value: v,
	}
}

func NewInt(v int) Int {
	return Int{
		Set: true,
		Value: v,
	}
}

func NewInt64(v int64) Int64 {
	return Int64{
		Set:   true,
		Value: v,
	}
}

func (o *String) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	o.Value = v
	o.Set = true
	return nil
}

func (o *Int) UnmarshalJSON(b []byte) error {
	var v int
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	o.Value = v
	o.Set = true
	return nil
}

func (o *Int64) UnmarshalJSON(b []byte) error {
	var v int64
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	o.Value = v
	o.Set = true
	return nil
}