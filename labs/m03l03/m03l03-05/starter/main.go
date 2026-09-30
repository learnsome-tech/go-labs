package main

import "fmt"

func work() {
	defer fmt.Println("cleanup")
	defer func() {
		if recover() != nil {
			fmt.Println("recovered")
		}
	}()
	panic("failure")
}

func main() {
	work()
	fmt.Println("done")
}
