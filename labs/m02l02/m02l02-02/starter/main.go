package main

import "fmt"

func main() {
	items := []string{"api", "worker"}
	items = append(items, "scheduler")
	fmt.Println(len(items))
	fmt.Println(items[1])
	fmt.Println(items[:2])
}
