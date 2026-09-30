package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan string)
	go func() { time.Sleep(time.Millisecond); result <- "ready" }()
	select {
	case value := <-result:
		fmt.Println(value)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}
