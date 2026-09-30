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
