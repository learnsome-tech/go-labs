package main

import "fmt"

func main() {
	var names []string
	names = append(names, "api", "worker")
	seen := make(map[string]bool)
	for _, name := range names {
		seen[name] = true
	}
	fmt.Println(names)
	fmt.Println(seen["api"], seen["cache"])
}
