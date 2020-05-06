package validator

import (
	"reflect"
	"testing"
)

type S []byte

func TestType(t *testing.T) {
	var s S
	var i interface{}
	var b []byte
	var a []byte = []byte("haha")
	var u []uint8

	i = s
	switch i.(type) {
	case []byte:
		t.Log("[]byte")
	case S:
		t.Log("S")
	}

	t.Log(reflect.ValueOf(i).Type().Kind())
	t.Log(reflect.TypeOf(s).Elem().String())
	t.Log(reflect.ValueOf(b).Kind())
	t.Log(reflect.ValueOf(b).String())
	t.Log(reflect.ValueOf(a).String())
	t.Logf("%s", reflect.ValueOf(a).Bytes())
	t.Logf("%s", reflect.ValueOf(u).Bytes())
}
