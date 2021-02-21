package main

import "fmt"

func main() {
	arr := [4]string{"shubh", "shahi", "shubham", "thaur"}

	fmt.Println("elements of an array are")

	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])
	}
}
