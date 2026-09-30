package main

import "fmt"

func divide(total, count int) (int, error) {
	if count == 0 {
		return 0, fmt.Errorf("count is zero")
	}
	return total / count, nil
}

func main() {
	value, err := divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("result:", value)
}
