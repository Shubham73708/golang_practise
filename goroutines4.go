package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	elapsedTime := time.Since(start)

	fmt.Println("total time in execution is:" + elapsedTime.String())
	time.Sleep(time.Second)

}
