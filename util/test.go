package util

import "testing"

func AssertEquals(t *testing.T, a, b interface{}) bool {
	if a == b {
		return true
	} else {
		t.Errorf("Expected %v, got %v", a, b)
		return false
	}
}
