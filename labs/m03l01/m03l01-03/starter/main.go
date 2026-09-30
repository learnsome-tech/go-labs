package main

import "fmt"

func load(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name is empty")
	}
	return "loaded " + name, nil
}

func main() {
	value, err := load("")
	if err != nil {
		fmt.Println("load failed:", err)
		return
	}
	fmt.Println(value)
}
