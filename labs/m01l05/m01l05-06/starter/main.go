package main

import "fmt"

func main() {
	var total int = 200
	var used int = 164
	var factor float64 = 1.25

	fmt.Println(float64(used) / float64(total) * factor)
	fmt.Println(used * int(factor))
	fmt.Println(float64(total) * factor)
}
