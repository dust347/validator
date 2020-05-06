package validator_test

import (
	"reflect"
	"testing"
	"validator"
)

func TestCheckValid(t *testing.T) {
	s := make([]string, 12)
	if err := validator.CheckValid(s, "length=12"); err != nil {
		t.Fatal(err)
	}

	if err := validator.CheckValid(s, "length=2"); err != nil {
		t.Log(err.Error())
	} else {
		t.FailNow()
	}
}

//
func TestUint2Int(t *testing.T) {
	var u uint64 = 64
	var i interface{}
	i = u

	in, ok := i.(int64)
	if ok {
		t.Log("ok")
	} else {
		t.Log("failed")
	}

	t.Logf("in type = %s", reflect.ValueOf(in).Type().String())
	t.Log(in)
	t.Log(int64(reflect.ValueOf(i).Uint()))
	t.Log(reflect.ValueOf(1).Type().String())
}
