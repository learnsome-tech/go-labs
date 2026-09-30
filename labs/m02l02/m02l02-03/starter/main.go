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
