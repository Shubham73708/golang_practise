package main

import "fmt"

func main() {
	slice := []string{"shubh", "shubham", "shahi", "ram", "shyaam"}

	for index, ele := range slice {
		fmt.Printf("index = %d, element=%s\n", index+2, ele)
	}

	for _, ele := range slice {
		fmt.Printf("element=%s\n", ele)
	}

}
