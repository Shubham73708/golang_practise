package main

import (
	"fmt"
	"time"
)

func f(str string) {
	for i := 0; i < 6; i++ {
		fmt.Println(str, ":", i)
	}
}
func main() {

	f("shubham")

	go f("shahi")

	go func(msg string) {
		fmt.Println(msg)
	}("shahi")

	time.Sleep(time.Second)
	fmt.Println("shuuu")
}
