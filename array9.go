package main

import "fmt"

func main() {

	array := [...]int{1, 2, 4, 5, 6, 7, 8}
	array1 := array
	fmt.Println("elements of an array are:", array)

	array1[0] = 8
	fmt.Println("elements of an array1 are:", array1)
}
