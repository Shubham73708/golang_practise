package main

import "fmt"

func main() {

	//reminder: In Go language, you can also concatenate string using Sprintf() method.

	s1 := "shubham "
	s2 := "shahi "
	s3 := "shubh"

	result := fmt.Sprintf("%s%s%s", s1, s2, s3)

	fmt.Println("result is:", result)

	//reminder: += operator adds a new or given string to the end of the specified string.
	s1 += s2
	fmt.Println(s1)
}
