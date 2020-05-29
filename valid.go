package validator

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

/*
//CheckStruct 检查结构体的值是否合法
func CheckStruct(v interface{}) error {
	return checkValue(reflect.ValueOf(v))
}

func checkValue(v reflect.Value) error {
	kind := v.Kind()

	//指针
	if (kind == reflect.Ptr || kind == reflect.Interface) && !v.IsNil() {
		return checkValue(v.Elem())
	}

	//不是结构体
	if kind != reflect.Struct {
		return errors.New("not struct")
	}

	return nil
}
*/
