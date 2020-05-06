package validator_test

import (
	"testing"
	"validator"
)

func TestCheckLen(t *testing.T) {
	s := make([]string, 12)

	r := validator.NewLengthRule(5, 13)
	err := r.Check(s)
	if err != nil {
		t.Fatal(err.Error())
	}

	s = make([]string, 4)
	err = r.Check(s)
	if err == nil {
		t.Fail()
	} else {
		t.Log(err.Error())
	}

	s = make([]string, 20)
	err = r.Check(s)
	if err == nil {
		t.Fail()
	} else {
		t.Log(err.Error())
	}
}

func TestCheckLenEqual(t *testing.T) {
	s := make([]string, 2)

	r := validator.NewLengthEqualRule(2)
	err := r.Check(s)
	if err != nil {
		t.Fatal(err.Error())
	}

	s = make([]string, 1)
	err = r.Check(s)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Fail()
	}

	s = make([]string, 3)
	err = r.Check(s)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Fail()
	}
}

func TestCheckLenMax(t *testing.T) {
	s := make([]string, 2)

	r := validator.NewLengthMaxRule(2)
	err := r.Check(s)
	if err != nil {
		t.Fatal(err.Error())
	}

	s = make([]string, 1)
	err = r.Check(s)
	if err != nil {
		t.Fatal(err.Error())
	}

	s = make([]string, 3)
	err = r.Check(s)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Fail()
	}
}

func TestCheckLenMin(t *testing.T) {
	s := make([]string, 2)

	r := validator.NewLengthMinRule(3)
	err := r.Check(s)
	if err == nil {
		t.Fatal("fail")
	} else {
		t.Log(err.Error())
	}

	s = make([]string, 3)
	if err = r.Check(s); err != nil {
		t.Fatal(err.Error())
	}

	s = make([]string, 4)
	if err = r.Check(s); err != nil {
		t.Fatal(err.Error())
	}
}

func TestMaxRule(t *testing.T) {
	//int
	ir := validator.NewMaxRule(int64(6))

	var err error
	if err = ir.Check(0); err != nil {
		t.Fatal(err)
	}

	if err = ir.Check(8); err != nil {
		t.Log(err)
	} else {
		t.Fatal("check failed")
	}

	if err = ir.Check("0"); err != nil {
		t.Log(err)
	} else {
		t.Fatal("fail")
	}

	if err = ir.Check(nil); err != nil {
		t.Log(err)
	} else {
		t.Fatal("fail")
	}

	if err = ir.Check(10.0); err != nil {
		t.Log(err)
	} else {
		t.Fatal("fail")
	}

	if err = ir.Check(uint(2)); err != nil {
		t.Fatal(err)
	}

	if err = ir.Check(uint(20)); err != nil {
		t.Log(err)
	}
}
