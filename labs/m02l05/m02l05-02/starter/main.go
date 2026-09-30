package main

import "fmt"

func show(value any) {
	text, ok := value.(string)
	if !ok {
		fmt.Println("not text")
		return
	}
	fmt.Println("text:", text)
}

func main() {
	show("ready")
	show(42)
}
