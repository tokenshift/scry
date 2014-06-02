package parser

import "testing"

func TestLit(t *testing.T) {
	v := StringVessel("2")
	res, err := Lit("2")(v)
	if nil != err {
		t.Error(err)
		return
	}
	if "2" != res {
		t.Errorf("Expected 2, got %s", res)
	}

	res, err = Lit("3")(v.Reset())
	if nil == err {
		t.Errorf("Expected failure, got %s", res)
	}
}

func TestOr(t *testing.T) {
	v := StringVessel("2")
	res, err := Or(Lit("2"), Lit("3"))(v)
	if nil != err {
		t.Error(err)
		return
	}
	if "2" != res {
		t.Errorf("Expected 2, got %s", res)
	}

	v = StringVessel("3")
	res, err = Or(Lit("2"), Lit("3"))(v)
	if nil != err {
		t.Error(err)
		return
	}
	if "3" != res {
		t.Errorf("Expected 3, got %s", res)
	}

	v = StringVessel("4")
	res, err = Or(Lit("2"), Lit("3"))(v)
	if nil == err {
		t.Errorf("Expected failure, got %s", res)
	}
}
