# Concurrent Systems Programming with Go — lesson m01l04 — Your First Program and go run
# https://learnsome.tech/courses/go-course/watch?lesson=m01l04
# © LearnSome.tech
package main

import (
	"errors"
	"fmt"
)

func usage(total, used int) (int, error) {
	if total <= 0 {
		return 0, errors.New("total must be positive")
	}
	return used * 100 / total, nil
}

func main() {
	pct, err := usage(200, 164)
	fmt.Println(pct, err)

	pct, err = usage(0, 164)
	fmt.Println(pct, err)
}
