package main

import "fmt"

func main() {

	array := [...]int{1, 2, 4, 5, 6, 7, 8}

	for x := 0; x < len(array); x++ {
		fmt.Printf("Elements of an array are %d\n", array[x])
	}
}
