package main

import (
	"fmt"
)

func myfunc(ch chan int) {
	fmt.Println(778 + <-ch)
}
func main() {

	fmt.Println("start")

	ch := make(chan int)

	go myfunc(ch)
	ch <- 45
	fmt.Println("end")
}
