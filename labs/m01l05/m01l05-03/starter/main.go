package main

import "fmt"

type limits struct {
	Retries int
	Timeout string
	Strict  bool
}

func main() {
	var count int
	var name string
	var ok bool
	var hosts []string
	var tags map[string]string
	var l limits

	fmt.Printf("%d %q %t\n", count, name, ok)
	fmt.Println(hosts == nil, len(hosts), tags == nil, len(tags))
	fmt.Printf("%+v\n", l)
}
