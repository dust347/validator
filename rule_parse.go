package validator

import (
	"fmt"
	"strings"
	"sync"
)

//rule name
const (
	//length
	ruleLength    = "length" //e.g. length=2
	ruleMaxLength = "max-length"
	ruleMinLength = "min-length"

	//max-min
	ruleMax = "max" //e.g. max=2
	ruleMin = "min"
)

//ParseRuleFn 解析单个规则，将string的规则描述解析为Rule实例
//其中 param 是等号后面的字符串
//e.g. rule=2 param 是 "2"
type ParseRuleFn func(param string) (Rule, error)

//valide rule parse
var validRuleParseFn map[string]ParseRuleFn = map[string]ParseRuleFn{
	//default rule parse

	//length
	ruleLength:    parseLengthRule,
	ruleMaxLength: parseMaxLengthRule,
	ruleMinLength: parseMinLengthRule,

	//max-min
	ruleMax: parseMaxRule,
	ruleMin: parseMinRule,
}

var mux sync.RWMutex //validRuleParseFn 的锁，防止并发注册

//RegisterRuleParseFn 注册一个新的解析函数
func RegisterRuleParseFn(ruleName string, parseFn ParseRuleFn) {
	mux.Lock()
	validRuleParseFn[ruleName] = parseFn
	mux.Unlock()
}

func getParseFn(ruleName string) ParseRuleFn {
	mux.RLock()
	r := validRuleParseFn[ruleName]
	mux.RUnlock()

	return r
}

func parseRule(str string) (Rule, error) {
	if str == "" {
		return emptyRule{}, nil
	}
	r := strings.Split(str, "=")
	if len(r) == 0 {
		return nil, fmt.Errorf("%s: rule parse failed", r)
	}

	/*
		parseFn, ok := validRuleParseFn[r[0]]
		if !ok {
			return nil, fmt.Errorf("%s: have no valid parse fn", r)
		}
	*/
	parseFn := getParseFn(r[0])
	if parseFn == nil {
		return nil, fmt.Errorf("%s: have no valid parse fn", r)
	}

	var param string
	if len(r) > 1 {
		param = r[1]
	}

	rule, err := parseFn(param)
	if err != nil {
		return nil, fmt.Errorf("%s: parse err, %s", r, err.Error())
	}

	return rule, nil
}
