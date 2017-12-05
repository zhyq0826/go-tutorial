package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hello Go!"))
	fmt.Printf("%T\n", w)
	rw := w.(io.ReadWriter)
	fmt.Printf("%T\n", rw)
}
