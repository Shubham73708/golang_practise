package main

import (
	"bytes"
	"fmt"
)

func main() {
	//reminder bytes.buffer is used to concatenate the string.
	//buffer is a temporary region of volatile memory (RAM) where data can be stored before consuming it.
	var b bytes.Buffer

	b.WriteString("s")
	b.WriteString("h")
	b.WriteString("u")
	b.WriteString("b")
	b.WriteString("h")
	b.WriteString("a")
	b.WriteString("m")

	fmt.Println("string is :", b.String())

}
