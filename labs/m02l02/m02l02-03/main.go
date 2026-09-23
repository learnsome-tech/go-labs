# Concurrent Systems Programming with Go — lesson m02l02 — Arrays, Slices, and Maps
# https://learnsome.tech/courses/go-course/watch?lesson=m02l02
# © LearnSome.tech
package main

import "fmt"

func main() {
	ports := map[string]int{"http": 80, "https": 443}
	port, ok := ports["ssh"]
	if !ok {
		fmt.Println("ssh is absent")
		return
	}
	fmt.Println(port)
}
