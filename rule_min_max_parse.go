package validator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func parseMaxRule(param string) (Rule, error) {
	if param == "" {
		return nil, errors.New("parse max rule failed, param empty")
	}

	max, err := parseCompareValue(param)
	if err != nil {
		return nil, fmt.Errorf("parse max value err, %s", err.Error())
	}

	return NewMaxRule(max), nil
}

func parseMinRule(param string) (Rule, error) {
	if param == "" {
		return nil, errors.New("parse min rule failed, param empty")
	}

	min, err := parseCompareValue(param)
	if err != nil {
		return nil, fmt.Errorf("parse min value err, %s", err.Error())
	}

	return NewMinRule(min), nil
}

//从字符串里面解析要比较的值 int64 or float64
func parseCompareValue(str string) (interface{}, error) {
	if strings.ContainsRune(str, '.') { //float
		return strconv.ParseFloat(str, 64)
	}

	//int
	return strconv.ParseInt(str, 10, 64)
}
