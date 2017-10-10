package main

import (
	"fmt"
)

type ByteCounter int

//func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func main() {

	var c ByteCounter
	c.Write([]byte("hello"))
	//Println formats using the default formats for its operands and writes to standard output
	fmt.Println(c)
	fmt.Fprintf(&c, "hello")
	fmt.Println(c)
}

func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)
func Printf(format string, args ...interface{}) (int, error) {
	return Fprintf(os.Stdout, format, args...)
}
func Sprintf(format string, args ...interface{}) string {
	var buf bytes.Buffer
	Fprintf(&buf, format, args...)
	return buf.String()
}
