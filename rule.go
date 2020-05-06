package validator

//Rule 单条校验规则
type Rule interface {
	Check(v interface{}) error
}

type emptyRule struct {
}

func (r emptyRule) Check(v interface{}) error {
	return nil
}
