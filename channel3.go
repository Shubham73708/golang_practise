package main

import "fmt"

func myfun(mychan chan string) {
	for s := 0; s < 6; s++ {
		mychan <- "shubham shahi"
	}
	close(mychan)
}

func main() {
	c := make(chan string)

	go myfun(c)

	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("channel is closed", ok)
			break
		}
		fmt.Println("channel is open", res, ok)
	}
}
