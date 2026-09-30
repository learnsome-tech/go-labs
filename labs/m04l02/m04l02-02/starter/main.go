package main

import "fmt"

func main() {
	values := make(chan string)
	go func() { values <- "ready" }()
	value := <-values
	fmt.Println(value)
}
