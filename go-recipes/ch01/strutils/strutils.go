package strutils

import (
	"strings"
	"unicode"
)

func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

func ToLowerCase(s string) string {
	return strings.ToLower(s)
}

func ToFirstUpper(s string) string {
	if len(s) < 1 {
		return s
	}

	t := strings.Trim(s, " ")
	t = strings.ToLower(t)
	res := []rune(t)

	res[0] = unicode.ToUpper(res[0])

	return string(res)
}
