package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(time.Millisecond)
	fmt.Println(time.Since(start) > 0)
}
