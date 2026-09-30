package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println("args:", args)
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: probe <action> <host>")
		os.Exit(2)
	}
	fmt.Println("action:", args[0], "host:", args[1])
	fmt.Println("timeout:", os.Getenv("PROBE_TIMEOUT"))
	if v, ok := os.LookupEnv("PROBE_DEBUG"); ok {
		fmt.Println("debug:", v)
	} else {
		fmt.Println("debug: unset")
	}
}
