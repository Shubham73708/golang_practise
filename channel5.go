package main

import "fmt"

func main() {
	mychan := make(chan string, 5)

	mychan <- "s"
	mychan <- "s"
	mychan <- "s"
	mychan <- "s"
	mychan <- "s"

	fmt.Println("length of channel is :", len(mychan))
	fmt.Println("capacity of channel is :", cap(mychan))

}
