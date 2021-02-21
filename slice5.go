package main

import "fmt"

func main() {
	slice1 := []string{"shubh", "shubham", "shahi", "ram", "shyaam"}

	for x := 0; x < len(slice1); x++ {
		fmt.Println("value of slice1 are:", slice1[x])
	}

}
