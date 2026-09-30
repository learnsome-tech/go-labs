package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "address to probe")
	tries := flag.Int("tries", 3, "attempts before giving up")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage: probe [flags] host...")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Println(*addr, *tries)
}
