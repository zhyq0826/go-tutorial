package pkgtwo

import (
	"fmt"

	"github.com/zhyq0826/go-tutorial/topic/test-package/db"
)

func Sum(a, b float64) float64 {
	fmt.Println("===> gloabl ", db.Global)
	return a + b
}
