package lib

import (
	"fmt"
)

func PrintfFavorites() {
	for _, v := range favorites {
		fmt.Println(v)
	}
}
