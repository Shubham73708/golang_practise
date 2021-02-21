package main

import "fmt"

func main() {
	arr := [5]string{"shubh", "shubham", "shahi", "ram", "shyaam"}

	fmt.Println("strings in array are:", arr)

	slice := arr[2:5]
	slice1 := arr[:5]
	slice2 := arr[2:]
	slice3 := arr[:]
	slice4 := arr[1:2]

	fmt.Println("strings in slice are:", slice)
	fmt.Println("strings in slice1 are:", slice1)
	fmt.Println("strings in slice2 are:", slice2)
	fmt.Println("strings in slice3 are:", slice3)
	fmt.Println("strings in slice4 are:", slice4)

}
