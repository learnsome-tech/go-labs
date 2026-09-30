package main

import "fmt"

func main() {
	values := make(chan int, 2)
	values <- 3
	values <- 5
	close(values)
	for value := range values {
		fmt.Println(value)
	}
	fmt.Println("done")
}
