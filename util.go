package validator

import (
	"fmt"
	"reflect"
)

const (
	intMax  = ^uint64(0) >> 1
	minUint = 0
)

func length(v interface{}) (int, error) {
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Slice, reflect.String, reflect.Map, reflect.Array:
		return value.Len(), nil
	}

	return 0, fmt.Errorf("can not get len of %s", value.Kind().String())
}

func asInt(v interface{}) (int64, error) {
	value := reflect.ValueOf(v)

	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int(), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr: //无符号整型
		if value.Uint() > intMax {
			return 0, fmt.Errorf("%d greater than max int, can not use as int", value.Uint())
		}
		return int64(value.Uint()), nil

	case reflect.Float32, reflect.Float64:
		return int64(value.Float()), nil
	}

	return 0, fmt.Errorf("%s can not use as int", value.Kind())
}

func asUint(v interface{}) (uint64, error) {
	value := reflect.ValueOf(v)

	switch value.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr: //无符号整型
		return value.Uint(), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if value.Int() < minUint {
			return 0, fmt.Errorf("%d less than min uint, can not use as uint", value.Int())
		}
		return uint64(value.Int()), nil

	case reflect.Float32, reflect.Float64:
		if value.Float() < minUint {
			return 0, fmt.Errorf("%f less than min uint, can not use as uint", value.Float())
		}
		return uint64(value.Float()), nil
	}

	return 0, fmt.Errorf("%s can not use as uint", value.Kind())
}

func asFloat(v interface{}) (float64, error) {
	value := reflect.ValueOf(v)

	switch value.Kind() {
	case reflect.Float32, reflect.Float64:
		return value.Float(), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(value.Int()), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr: //无符号整型
		return float64(value.Uint()), nil
	}

	return 0, fmt.Errorf("%s can not use as float", value.Kind())
}
