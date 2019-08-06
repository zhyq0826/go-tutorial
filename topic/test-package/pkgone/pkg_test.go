package pkgone

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Parallel()
	if Sum(1, 1) != 2 {
		t.Fatal("not equal")
	}
}

func TestSum2(t *testing.T) {
	// db.Global = 0
	t.Parallel()
	if Sum(1, 1) != 2 {
		t.Fatal("not equal")
	}
}

func TestMain(t *testing.M) {
	fmt.Println("hello main")
	t.Run()
}
