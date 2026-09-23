# Concurrent Systems Programming with Go — lesson m05l01 — Files and the os Package
# https://learnsome.tech/courses/go-course/watch?lesson=m05l01
# © LearnSome.tech
package main

import (
	"fmt"
	"os"
)

func main() {
	conf := []byte("interval=30s\nretries=3\n")
	err := os.WriteFile("probe.conf", conf, 0o600)
	if err != nil {
		fmt.Fprintln(os.Stderr, "write:", err)
		os.Exit(1)
	}
	back, err := os.ReadFile("probe.conf")
	if err != nil {
		fmt.Fprintln(os.Stderr, "read:", err)
		os.Exit(1)
	}
	fmt.Printf("read %d bytes\n", len(back))
	fmt.Print(string(back))
}
