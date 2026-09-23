# Concurrent Systems Programming with Go — lesson m05l01 — Files and the os Package
# https://learnsome.tech/courses/go-course/watch?lesson=m05l01
# © LearnSome.tech
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("app.log")
	if err != nil {
		fmt.Fprintln(os.Stderr, "open:", err)
		os.Exit(1)
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, "stat:", err)
		os.Exit(1)
	}
	fmt.Printf("%s: %d bytes, mode %s\n", info.Name(), info.Size(), info.Mode())
}
