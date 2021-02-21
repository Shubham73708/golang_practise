package main

import "fmt"

func main() {
	slice1 := make([]int, 4, 8)
	fmt.Printf("slice1:%v, \nlength:%d, \ncapacity:%d\n", slice1, len(slice1), cap(slice1))

	slice2 := make([]int, 8)
	fmt.Printf("slice2:%v, \nlength:%d, \ncapacity:%d\n", slice2, len(slice2), cap(slice2))

}
