package validator

import "errors"

func parseReqexpRule(param string) (Rule, error) {
	if param == "" {
		return nil, errors.New("parse regexp rule failed, param is empty")
	}

	return NewRegexpRule(param), nil
}
