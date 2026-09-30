package main

import (
	"context"
	"fmt"
)

func report(ctx context.Context) {
	if value := ctx.Value("request"); value != nil {
		fmt.Println(value)
	}
}

func main() {
	ctx := context.WithValue(context.Background(), "request", "probe")
	report(ctx)
}
