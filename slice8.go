package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := []int{10, 2, 30, 4, 50, 6, 70, 8, 90}
	s2 := []string{"shubh", "shahi", "shubham", "sonal", "suraj"}

	fmt.Println("Before sorting:")
	fmt.Println("Before sorting:", s1)
	fmt.Println("Before sorting:", s2)

	sort.Strings(s2)
	sort.Ints(s1)

	fmt.Println("After sorting:")
	fmt.Println("After sorting:", s1)
	fmt.Println("After sorting:", s2)
}
