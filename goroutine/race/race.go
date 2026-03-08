package main

import (
	"fmt"
	"time"
)

func main() {
	c := 1

	for i := 0; i < 1000; i++ {
		go func() {
			c++
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("Counter:", c)
}

// go run -race
