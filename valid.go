package validator

import (
	"reflect"
)

//Check 校验某个值是否合法
func Check(v interface{}, rules ...Rule) error {
	var err error
	for _, r := range rules {
		err = r.Check(v)
		if err != nil {
			return err
		}
	}

	return nil
}

//CheckValid 校验某个值是否合法
//rules为字符串形式的规则描述 e.g. length=2
func CheckValid(v interface{}, strRules ...string) error {
	var rules []Rule

	for _, str := range strRules {
		r, err := parseRule(str)
		if err != nil {
			return err
		}

		rules = append(rules, r)
	}

	return Check(v, rules...)
}

//CheckStruct 检查结构体的值是否合法
func CheckStruct(s interface{}) error {
	v := reflect.TypeOf(s)

	//不是结构体就直接返回
	if v.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < v.NumField(); i++ {

	}

	return nil
}
