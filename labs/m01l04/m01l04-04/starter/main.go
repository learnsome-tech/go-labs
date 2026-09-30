package main

import "fmt"

type target struct {
	Host string
	Port int
}

func main() {
	t := target{Host: "db-1", Port: 5432}
	fmt.Printf("host %s port %d\n", t.Host, t.Port)
	fmt.Printf("quoted %q ready %t\n", t.Host, true)
	fmt.Printf("value %v\n", t)
	fmt.Printf("fields %+v\n", t)
	fmt.Printf("type %T\n", t)
	line := fmt.Sprintf("%s:%d", t.Host, t.Port)
	fmt.Println(line)
}
