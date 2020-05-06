package validator

import (
	"errors"
	"fmt"
)

//LengthRule 长度相关的规则
type LengthRule struct {
	max int
	min int
}

//NewLengthRule 创建LengthRule实例
//min, max是长度最小值和最大值
//如果max < 0. min >= 0, 则length最小值为min
func NewLengthRule(min, max int) LengthRule {
	return LengthRule{
		min: min,
		max: max,
	}
}

//NewLengthMaxRule 创建LengthRule实例
//表示长度不能超过max
func NewLengthMaxRule(max int) LengthRule {
	return LengthRule{
		min: -1,
		max: max,
	}
}

//NewLengthMinRule 创建LengthRule实例
//表示长度最小为min
func NewLengthMinRule(min int) LengthRule {
	return LengthRule{
		min: min,
		max: -1,
	}
}

//NewLengthEqualRule 创建LengthRule实例
//表示长度必须等于l
func NewLengthEqualRule(l int) LengthRule {
	return LengthRule{
		min: l,
		max: l,
	}
}

//Check 校验某个值是否符合长度规则
func (r LengthRule) Check(v interface{}) error {
	if r.max < 0 && r.min < 0 {
		return errors.New("max and min is invalid, must >= 0")
	}

	if (r.max > 0 && r.min > 0) && r.max < r.min {
		return errors.New("max must more than min")
	}

	//获取len
	l, err := length(v)
	if err != nil {
		return err
	}

	//between
	if r.max >= 0 && r.min >= 0 {
		if r.min <= l && l <= r.max {
			return nil
		}

		if r.max == r.min {
			return fmt.Errorf("length must be equal to %d", r.min)
		}
		return fmt.Errorf("length must be between %d and %d", r.min, r.max)
	}

	//单个
	if r.max >= 0 && l > r.max {
		return fmt.Errorf("length must not exceed %d", r.max)
	}

	if r.min >= 0 && l < r.min {
		return fmt.Errorf("length must not be less than %d", r.min)
	}

	return nil
}
