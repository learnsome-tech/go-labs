package main

import "fmt"

func timeout(value any) (int, bool) {
	seconds, ok := value.(int)
	if !ok || seconds < 1 {
		return 0, false
	}
	return seconds, true
}

func main() {
	if seconds, ok := timeout(5); ok {
		fmt.Println("timeout", seconds)
	}
	if _, ok := timeout("five"); !ok {
		fmt.Println("invalid timeout")
	}
}
