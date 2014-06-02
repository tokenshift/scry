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

	v.Reset()
	res, err = Lit("3")(v)
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

func TestSeq(t *testing.T) {
	v := StringVessel("42")
	res, err := Seq(Lit("4"), Lit("2"))(v)
	if nil != err {
		t.Error(err)
		return
	}
	if !eq([]interface{} { "4", "2" }, res) {
		t.Errorf("Expected 42, got %v", res)
	}

	v.Reset()
	res, err = Seq(Lit("4"), Lit("4"))(v)
	if nil == err {
		t.Errorf("Expected failure, got %s", res)
	}
}

// Compares two results to see if they are equal.
func eq(a, b Result) bool {
	if as, ok := a.([]interface{}); ok {
		if bs, ok := b.([]interface{}); ok {
			if len(as) == len(bs) {
				for i := range(as) {
					if !eq(as[i], bs[i]) {
						return false
					}
				}
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else if a == b {
		return true
	} else {
		return false
	}
}
