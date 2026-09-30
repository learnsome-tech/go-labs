package main

import "fmt"

func runWorker() (err error) {
	defer func() {
		if value := recover(); value != nil {
			err = fmt.Errorf("worker panic: %v", value)
		}
	}()
	panic("bad invariant")
}

func main() {
	if err := runWorker(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("supervisor alive")
}
