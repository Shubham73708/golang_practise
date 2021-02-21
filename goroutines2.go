package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome ")

	go func() {
		fmt.Println("on my")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("page")

}
