# Concurrent Systems Programming with Go — lesson m05l01 — Files and the os Package
# https://learnsome.tech/courses/go-course/watch?lesson=m05l01
# © LearnSome.tech
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	for _, path := range []string{"probe.conf", "gone.conf"} {
		info, err := os.Stat(path)
		switch {
		case err == nil:
			fmt.Printf("%s: present, %d bytes\n", path, info.Size())
		case errors.Is(err, os.ErrNotExist):
			fmt.Printf("%s: absent\n", path)
		default:
			fmt.Printf("%s: unreadable: %v\n", path, err)
		}
	}
}
