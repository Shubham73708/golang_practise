package main

import "fmt"

func main() {

	x := 0xFF
	y := 0x9C

	fmt.Printf("value of x is %T\n", x)
	fmt.Printf("value of x in hexadecimal is %X\n", x)
	fmt.Printf("value of x in decimal is %v\n", x)

	fmt.Printf("value of y is %T\n", y)
	fmt.Printf("value of y in hexadecimal is %X\n", y)
	fmt.Printf("value of y in decimal is %v", y)

}
