package pkgone_test

import (
	"testing"

	. "github.com/zhyq0826/go-tutorial/topic/test-package/pkgone"
)

func TestSum3(t *testing.T) {
	t.Parallel()
	if Sum(1, 1) != 2 {
		t.Fatal("not equal")
	}
}
