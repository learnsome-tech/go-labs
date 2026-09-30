package main

import "fmt"

func describe(value *string) {
	if value == nil {
		fmt.Println("unset")
		return
	}
	fmt.Println("value:", *value)
}

func main() {
	var note *string
	describe(note)
	text := "ready"
	describe(&text)
}
