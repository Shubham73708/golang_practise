package main

import "fmt"

func main() {
	arr := [3][3]string{{"shubh", "shahi", "shubham"},
		{"shubham", "shahi", "shubh"},
		{"shahi", "shubham", "shubh"}}

	fmt.Println("element of array 1:")
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Println(arr[x][y])
		}
	}
}
