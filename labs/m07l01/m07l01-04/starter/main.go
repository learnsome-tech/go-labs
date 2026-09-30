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
