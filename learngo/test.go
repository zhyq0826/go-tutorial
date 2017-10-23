package main

import (
	"bytes"
	"fmt"
	"io"
)

type T struct {
}

func (t *T) String() string {
	return ""
}

func main() {
	var t T
	t.String()
	// T{}.String()
	var _ fmt.Stringer = &t
	// cannot use t (type T) as type fmt.Stringer in assignment: T does not implement fmt.Stringer (String method has pointer receiver)
	// var _ fmt.Stringer = t
	var _ io.Writer = new(bytes.Buffer)
	var _ io.Writer = (*bytes.Buffer)(nil)
}
