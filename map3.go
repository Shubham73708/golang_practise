package main

import (
	"fmt"
)

func main() {
	map3 := map[int]string{
		1: "shubh",
		2: "shubham",
		3: "shahi",
		4: "suraj",
		5: "sonal",
	}

	fmt.Println("original value of map is:", map3)

	map3[6] = "shaurya"
	map3[7] = "shubhi"

	fmt.Println("added value of map is:", map3)

	map3[2] = "surya"
	map3[3] = "shubhrastra"

	fmt.Println("updated value of map is:", map3)

	value1 := map3[3]
	value2 := map3[4]

	fmt.Println("value1 of map is:", value1)
	fmt.Println("value2 of map is:", value2)

	value3, ok := map3[5]
	fmt.Println("value3 of map is found or not:", ok)
	fmt.Println("value3 of map is:", value3)

	_, ok1 := map3[5]
	fmt.Println("value of map is found or not:", ok1)

	delete(map3, 5)
	delete(map3, 4)

	fmt.Println("deleted value of map is:", map3)

}
