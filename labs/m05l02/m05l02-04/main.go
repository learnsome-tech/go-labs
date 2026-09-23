# Concurrent Systems Programming with Go — lesson m05l02 — Streams and the io Package
# https://learnsome.tech/courses/go-course/watch?lesson=m05l02
# © LearnSome.tech
package main

import (
	"fmt"
	"io
	"strings
)

func main() {
	reader := io.LimitReader(strings.NewReader("0123456789"), 4)
	data, err := io.ReadAll(reader)
	fmt.Println(string(data), err)
}
