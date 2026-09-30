package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2025, time.January, 2, 3, 4, 5, 0, time.UTC)
	fmt.Println(t.Format(time.RFC3339))
}
