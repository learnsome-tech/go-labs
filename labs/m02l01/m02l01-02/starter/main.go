package main

import "fmt"

func main() {
	name := "api"
	p := &name
	fmt.Println(*p)
	*p = "worker"
	fmt.Println(name)
}
