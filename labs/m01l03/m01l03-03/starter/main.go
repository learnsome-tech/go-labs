package main

import (
	"fmt"

	"course.local/probe/internal/check"
)

func main() {
	fmt.Println(check.Describe(check.Target{Name: "db-1", Port: 5432}))
}
