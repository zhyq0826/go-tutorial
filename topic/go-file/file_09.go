package main

import (
	"bufio"
	"errors"
	"fmt"
)

type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
	fmt.Printf("Write: %q\n", p)
	return 0, errors.New("IO Error!")
}

func main() {
	wr := bufio.NewWriterSize(new(Writer), 3)
	wr.Write([]byte{'a'})
	wr.Write([]byte{'b'})
	wr.Write([]byte{'c'})
	wr.Write([]byte{'d'})
	err := wr.Flush()
	fmt.Println(err)
}
