# Concurrent Systems Programming with Go — lesson m07l01 — Project Setup and Configuration
# https://learnsome.tech/courses/go-course/watch?lesson=m07l01
# © LearnSome.tech
package main

import (
	"errors"
	"fmt"
)

type Target struct { Name, URL string }

func validate(t Target) error {
	if t.Name == "" || t.URL == "" {
		return errors.New("target needs name and URL")
	}
	return nil
}

func main() {
	err := validate(Target{Name: "api"})
	fmt.Println(err)
}
