package main

import "fmt"

func addTag(label *string) {
	*label = *label + "-ready"
}

func main() {
	status := "service"
	addTag(&status)
	fmt.Println(status)
}
