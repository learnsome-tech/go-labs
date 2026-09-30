package main

import "fmt"

type Reader interface {
	Read() string
}

type Fixture struct{}

func (Fixture) Read() string { return "fixture data" }

func load(r Reader) string {
	return r.Read()
}

func main() {
	fmt.Println(load(Fixture{}))
}
