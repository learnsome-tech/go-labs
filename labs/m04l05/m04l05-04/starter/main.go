package main

import (
	"context"
	"fmt"
)

func worker(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("worker stopped")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() { worker(ctx); close(done) }()
	cancel()
	<-done
}
