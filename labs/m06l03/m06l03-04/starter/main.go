package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	fmt.Println(err)
}
