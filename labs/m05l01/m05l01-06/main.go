# Concurrent Systems Programming with Go — lesson m05l01 — Files and the os Package
# https://learnsome.tech/courses/go-course/watch?lesson=m05l01
# © LearnSome.tech
package main

import (
	"fmt"
	"os"
)

const (
	exitOK      = 0
	exitFailed  = 1
	exitBadArgs = 2
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: probe <host>")
		os.Exit(exitBadArgs)
	}
	fmt.Println("checking", os.Args[1])
	fmt.Fprintln(os.Stderr, "probe: host unreachable")
	os.Exit(exitFailed)
}
