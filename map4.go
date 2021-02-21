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

	fmt.Println("original value of map3 is:", map3)
	map4 := map3
	map4[6] = "shaurya"
	map4[7] = "shubhi"

	fmt.Println("added value of map3 is:", map3)
	fmt.Println("added value of map4 is:", map4)

}
