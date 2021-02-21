package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 8}
	arr3 := [3]int{1, 2, 3}

	fmt.Println("length of array 1:", len(arr1))
	fmt.Println("length of array 2:", len(arr2))
	fmt.Println("length of array 3:", len(arr3))

	fmt.Println(arr1 == arr2)

	fmt.Println(arr1 == arr3)

	fmt.Println(arr2 == arr3)
}
