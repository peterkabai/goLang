package math

import "testing"

func TestAddition(t *testing.T) {
	v := Addition(5, 6)
	if v != 11 {
		t.Error("Expected 11, got ", v)
	}
}