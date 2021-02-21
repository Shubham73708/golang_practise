package main

import "fmt"

type Employee struct {
	firstname, lastname string
	age, salary         int
}

func main() {
	emp1 := &Employee{"shubh", "thakur", 24, 754656}

	fmt.Println("emplyoo info is:", *emp1)

	fmt.Println("emplyoo info is:", (*emp1).firstname)

	fmt.Println("emplyoo info is:", emp1.lastname)

}
