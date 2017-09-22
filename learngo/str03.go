package main

import "fmt"
import "unicode/utf8"

func main() {
	str := "Étoilé"
	rs := make([]byte, 0, 4)
	for i := 0; i < len(str); i++ {
		rs = append(rs, str[i])
		if utf8.FullRune(rs) {
			char, _ := utf8.DecodeRune(rs)
			fmt.Printf("%c", char)
			rs = rs[0:0]
		}
	}
}
