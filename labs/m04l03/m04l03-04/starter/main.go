package main

import "fmt"

func main() {
	updates := make(chan string, 1)
	select {
	case value := <-updates:
		fmt.Println(value)
	default:
		fmt.Println("no update")
	}
}
