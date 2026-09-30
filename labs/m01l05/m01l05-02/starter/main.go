package main

import "fmt"

var release = "stable"

func main() {
	host := "db-1"
	port := 5432
	ratio := 0.75
	var retries int64 = 3
	var enabled bool

	fmt.Printf("%T %T %T\n", host, port, ratio)
	fmt.Printf("%T %T %T\n", retries, enabled, release)
	fmt.Println(host, port, ratio, retries, enabled, release)
}
