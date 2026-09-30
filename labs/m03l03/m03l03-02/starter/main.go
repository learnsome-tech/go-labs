package main

import "fmt"

func lookup(items []string, index int) (string, error) {
	if index < 0 || index >= len(items) {
		return "", fmt.Errorf("index outside list")
	}
	return items[index], nil
}

func main() {
	value, err := lookup([]string{"api"}, 4)
	if err != nil {
		fmt.Println("lookup failed")
		return
	}
	fmt.Println(value)
}
