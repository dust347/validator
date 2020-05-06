package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

//RegexpRule 正则表达式相关的规则
type RegexpRule struct {
	reg          *regexp.Regexp
	customErrMsy string //自定义错误信息
}

func (r RegexpRule) getMatchFailErr() error {
	if r.customErrMsy != "" {
		return errors.New(r.customErrMsy)
	}

	return fmt.Errorf("can not match regexp %s", r.reg.String())
}

//Check 校验是否符合正则表达式
func (r RegexpRule) Check(v interface{}) error {
	value := reflect.ValueOf(v)

	switch value.Kind() {
	case reflect.String:
		if !r.reg.MatchString(value.String()) {
			return r.getMatchFailErr()
		}

	case reflect.Slice: //[]byte
		if value.Elem().Kind() != reflect.Uint8 {
			return errors.New("only string or []byte could be check")
		}

		if !r.reg.Match(value.Bytes()) {
			return r.getMatchFailErr()
		}

	default:
		return errors.New("only string or []byte could be check")
	}

	return nil
}
