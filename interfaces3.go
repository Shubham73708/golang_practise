package main

import "fmt"

func myfu(a interface{}) {
	value, ok := a.(float64)
	fmt.Println(value, ok)
}

func main() {
	var a1 interface{} = 96.43
	myfu(a1)
	var a2 interface{} = "shubham"
	myfu(a2)

}
