# Concurrent Systems Programming with Go — lesson m03l01 — Errors as Values
# https://learnsome.tech/courses/go-course/watch?lesson=m03l01
# © LearnSome.tech
package main

import "fmt"

func check(host string) error {
	if host != "web01" {
		return fmt.Errorf("host %s is unknown", host)
	}
	return nil
}

func main() {
	if err := check("cache01"); err != nil {
		fmt.Println("probe failed:", err)
		return
	}
	fmt.Println("probe passed")
}
