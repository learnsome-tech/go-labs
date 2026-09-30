package main

import (
	"flag"
	"fmt"
	"os"
)

func envOr(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}

func main() {
	addr := flag.String("addr", envOr("PROBE_ADDR", "127.0.0.1:8080"), "address")
	flag.Parse()
	fmt.Println("addr:", *addr)
}
