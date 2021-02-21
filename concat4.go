package main

import (
	"fmt"
	"strings"
)

func main() {

	slice := []string{"shubham", "shahi", "shubh"}

	result := strings.Join(slice, "-")
	fmt.Println(result)

}
