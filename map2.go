package main

import "fmt"

func main() {
	var map1 = make(map[float64]string)

	fmt.Println("map1 value is :", map1)

	map1[1.20] = "shubh"
	map1[1.30] = "shahi"
	fmt.Println("map1 value is :", map1)

}
