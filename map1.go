package main

import "fmt"

func main() {
	var map1 map[int]int

	if map1 == nil {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	map2 := map[int]string{
		1: "shubh",
		2: "shubham",
		3: "shahi",
		4: "suraj",
		5: "sonal",
	}
	fmt.Println("map2 are:", map2)

	for id, name := range map2 {
		fmt.Println(id, name)
	}
}
