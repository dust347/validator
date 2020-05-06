package validator

import (
	"errors"
	"fmt"
	"reflect"
)

type comparer interface {
	compareInt(int64, int64) error
	compareUint(uint64, uint64) error
	compareFloat(float64, float64) error
}

//大于等于
type greaterEqual struct{}

func (c greaterEqual) compareInt(l, r int64) error {
	if l >= r {
		return nil
	}

	return fmt.Errorf("must greater or equal %d", r)
}

func (c greaterEqual) compareUint(l, r uint64) error {
	if l >= r {
		return nil
	}

	return fmt.Errorf("must greater or equal %d", r)
}

func (c greaterEqual) compareFloat(l, r float64) error {
	if l >= r {
		return nil
	}

	return fmt.Errorf("must greater or equal %f", r)
}

type lessEqual struct{}

func (c lessEqual) compareInt(l, r int64) error {
	if l <= r {
		return nil
	}

	return fmt.Errorf("must less or equal %d", r)
}

func (c lessEqual) compareUint(l, r uint64) error {
	if l <= r {
		return nil
	}

	return fmt.Errorf("must less or equal %d", r)
}

func (c lessEqual) compareFloat(l, r float64) error {
	if l <= r {
		return nil
	}

	return fmt.Errorf("must less or equal %f", r)
}

//CompareRule min max之类的比较规则
type CompareRule struct {
	v interface{} //用于比较的值
	c comparer    //比较类型
}

//NewMaxRule 创建CompareRule实例
//表示必须小于等于max
func NewMaxRule(max interface{}) CompareRule {
	return CompareRule{
		v: max,
		c: lessEqual{},
	}
}

//NewMinRule 创建CompareRule实例
//表示必须大于等于min
func NewMinRule(min interface{}) CompareRule {
	return CompareRule{
		v: min,
		c: greaterEqual{},
	}
}

//Check 判断某个值是否满足比较规则
func (r CompareRule) Check(v interface{}) error {
	if !reflect.ValueOf(r.v).IsValid() {
		return errors.New("value to compare invalid")
	}

	value := reflect.ValueOf(v)
	if !value.IsValid() {
		return errors.New("value to check invalid")
	}

	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64: //整型
		rv, err := asInt(r.v)
		if err != nil {
			return fmt.Errorf("value to check is %s, can not compare, %s", value.Kind().String(), err.Error())
		}
		return r.c.compareInt(value.Int(), rv)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr: //无符号整型
		rv, err := asUint(r.v)
		if err != nil {
			return fmt.Errorf("value to check is %s, can not compare, %s", value.Kind().String(), err.Error())
		}
		return r.c.compareUint(value.Uint(), rv)

	case reflect.Float32, reflect.Float64:
		rv, err := asFloat(r.v)
		if err != nil {
			return fmt.Errorf("value to check is %s, can not compare, %s", value.Kind().String(), err.Error())
		}
		return r.c.compareFloat(value.Float(), rv)
	}

	return fmt.Errorf("type %s not support to compare", value.Type().String())
}
