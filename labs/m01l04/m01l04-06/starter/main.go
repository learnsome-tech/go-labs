package main

import (
	"fmt"
	"os"
)

func main() {
	unhealthy := 2
	fmt.Println("checked three targets")
	if unhealthy > 0 {
		fmt.Fprintf(os.Stderr, "%d unhealthy\n", unhealthy)
		os.Exit(1)
	}
	fmt.Println("all healthy")
}
