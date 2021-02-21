package main

import "fmt"

func main() {
	mychan := make(chan string)

	go func() {
		mychan <- "shubham"
		mychan <- "shahi"
		mychan <- "shubh"
		close(mychan)

	}()

	for res := range mychan {
		fmt.Println(res)
	}
}
