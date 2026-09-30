package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	tries := flag.Int("tries", 3, "attempts before giving up")
	flag.Parse()
	hosts := flag.Args()
	if len(hosts) == 0 {
		fmt.Fprintln(os.Stderr, "probe: no hosts given")
		os.Exit(2)
	}
	fmt.Println("hosts:", len(hosts), "tries:", *tries)
	for i, h := range hosts {
		fmt.Printf("%d %s\n", i+1, h)
	}
}
