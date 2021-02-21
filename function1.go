package main

import (
	"fmt"
)

func modify(x int) {
	x = 5
}

func main() {
	x := 10

	fmt.Println("value of x is:", x)

	modify(x)

	fmt.Println("value of x is:", x)
}
