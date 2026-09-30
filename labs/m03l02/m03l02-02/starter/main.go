package main

import (
	"errors"
	"fmt"
)

var ErrMissing = errors.New("missing")

func load() error {
	return fmt.Errorf("read config: %w", ErrMissing)
}

func main() {
	err := load()
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrMissing))
}
