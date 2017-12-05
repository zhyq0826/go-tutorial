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
	fw := w.(*os.File)
	fmt.Printf("%T\n", fw)
}
