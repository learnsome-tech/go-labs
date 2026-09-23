# Concurrent Systems Programming with Go — lesson m05l02 — Streams and the io Package
# https://learnsome.tech/courses/go-course/watch?lesson=m05l02
# © LearnSome.tech
package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	source := bytes.NewBufferString("logs ready")
	var target bytes.Buffer
	n, err := io.Copy(&target, source)
	fmt.Println(n, err, target.String())
}
