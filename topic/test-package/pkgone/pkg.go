package pkgone

import (
	"fmt"

	"github.com/zhyq0826/go-tutorial/topic/test-package/db"
)

func Sum(a, b int) int {
	db.Global += 1
	fmt.Println("===> global ", db.Global)
	return a + b
}


