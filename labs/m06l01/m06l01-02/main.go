# Concurrent Systems Programming with Go — lesson m06l01 — Parsing Arguments with flag
# https://learnsome.tech/courses/go-course/watch?lesson=m06l01
# © LearnSome.tech
package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "address to probe")
	tries := flag.Int("tries", 3, "attempts before giving up")
	verbose := flag.Bool("v", false, "log every attempt")
	wait := flag.Duration("timeout", 2*time.Second, "deadline per attempt")
	flag.Parse()
	fmt.Println("addr:", *addr)
	fmt.Println("tries:", *tries)
	fmt.Println("verbose:", *verbose)
	fmt.Println("timeout:", *wait)
}
