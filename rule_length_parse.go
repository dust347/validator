package validator

import (
	"errors"
	"fmt"
	"strconv"
)

func parseLengthRule(param string) (Rule, error) {
	if param == "" {
		return nil, errors.New("parse length rule failed, param is empty")
	}

	l, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf("parse length rule failed, %s", err.Error())
	}
	return NewLengthEqualRule(l), nil
}

func parseMaxLengthRule(param string) (Rule, error) {
	if param == "" {
		return nil, errors.New("parse max-length rule failed, param is empty")
	}

	l, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf("parse max-length rule failed, %s", err.Error())
	}

	return NewLengthMaxRule(l), nil
}

func parseMinLengthRule(param string) (Rule, error) {
	if param == "" {
		return nil, errors.New("parse min-length rule failed, param is empty")
	}

	l, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf("parse min-length rule failed, %s", err.Error())
	}

	return NewLengthMinRule(l), nil
}
