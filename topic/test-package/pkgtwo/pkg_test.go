package pkgtwo

import "testing"

func TestSum(t *testing.T) {
	if Sum(1, 1) != 2.0 {
		t.Fatal("not equal")
	}
}
