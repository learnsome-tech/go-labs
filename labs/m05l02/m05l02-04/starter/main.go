package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := io.LimitReader(strings.NewReader("0123456789"), 4)
	data, err := io.ReadAll(reader)
	fmt.Println(string(data), err)
}
