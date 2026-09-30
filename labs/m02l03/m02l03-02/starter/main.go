package main

import "fmt"

type Target struct {
	Name string
	Port int
}

func main() {
	target := Target{Name: "api", Port: 8080}
	fmt.Println(target.Name, target.Port)
}
