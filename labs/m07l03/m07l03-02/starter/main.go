package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := ctx.Err(); err != nil {
		fmt.Println("unhealthy:", err)
	}
}
