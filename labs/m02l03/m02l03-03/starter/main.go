package main

import "fmt"

type Target struct {
	Name string
	Port int
}

func (t Target) Address() string {
	return fmt.Sprintf("%s:%d", t.Name, t.Port)
}

func main() {
	t := Target{Name: "api", Port: 8080}
	fmt.Println(t.Address())
}
