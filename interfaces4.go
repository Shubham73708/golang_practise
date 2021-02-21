package main

import "fmt"

func myf(a interface{}) {
	switch a.(type) {

	case int:
		fmt.Println("integer value is:", a.(int))

	case float64:
		fmt.Println("integer value is:", a.(float64))

	case string:
		fmt.Println("integer value is:", a.(string))
	default:
		fmt.Println("there is default value")
	}

}

func main() {
	myf("shubh")
	myf(12)
	myf(true)
}
