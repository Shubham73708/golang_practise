package main

import "fmt"

func main() {

	var x int = 5478

	var p *int

	p = &x

	fmt.Println("value of x stored in:", x)
	fmt.Println("address of x:", &x)
	fmt.Println("value of variable p:", p)
	fmt.Println("value of variable p:", *p)
	*p = 500
	fmt.Println("value of variable p:", *p)

	var s *int

	fmt.Println("s:", s)
}
