package main

import (
	"fmt"
)

func main() {
	var shubh chan int
	fmt.Println("channel1 is :", shubh)
	fmt.Printf("type1 is : %T\n", shubh)

	shubham := make(chan int)
	fmt.Println("channel2 is :", shubham)
	fmt.Printf("type2 is : %T", shubham)
}
