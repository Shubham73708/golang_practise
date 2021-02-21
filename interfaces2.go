package main

import "fmt"

func myval(a interface{}) {
	val := a.(string)
	fmt.Println("value is:", val)
}

func main() {
	var val interface{} = "shubham shahi"
	myval(val)
}
