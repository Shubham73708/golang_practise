package main

import "fmt"

func main() {
	arr := [5]string{"shubh", "shubham", "shahi", "ram", "shyaam"}

	fmt.Println("strings in array are:", arr)

	slice := arr[1:5]

	fmt.Println("strings in slice are:", slice)

	fmt.Println("length of slice are:", len(slice))

	fmt.Println("capacity of slice are:", cap(slice))

}
