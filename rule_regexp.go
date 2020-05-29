package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

//RegexpRule 正则表达式相关的规则
type RegexpRule struct {
	pattern      string //正则
	customErrMsy string //自定义错误信息
}

//NewRegexpRule 创建 RegexpRule 实例
func NewRegexpRule(pattern string, customErrMsy ...string) *RegexpRule {
	return &RegexpRule{
		pattern:      pattern,
		customErrMsy: strings.Join(customErrMsy, ", "),
	}
}

//正则匹配失败生成错误文案
func (r *RegexpRule) getMatchErr(matched bool, err error) error {
	if err != nil { //正则解析失败
		return err
	}

	if matched { //正则匹配成功
		return nil
	}
	if r.customErrMsy != "" {
		return errors.New(r.customErrMsy)
	}

	return fmt.Errorf("can not match regexp %s", r.pattern)
}

//Check 校验是否符合正则表达式
func (r *RegexpRule) Check(v interface{}) error {
	value := reflect.ValueOf(v)

	switch value.Kind() {
	case reflect.String:
		return r.getMatchErr(regexp.MatchString(r.pattern, value.String()))

	case reflect.Slice: //[]byte
		if value.Elem().Kind() != reflect.Uint8 {
			return errors.New("only string or []byte could be check")
		}

		return r.getMatchErr(regexp.Match(r.pattern, value.Bytes()))

	default:
		return errors.New("only string or []byte could be check")
	}
	//return nil
}
