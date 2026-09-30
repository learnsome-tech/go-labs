package main

import (
	"errors"
	"fmt"
)

var ErrMissing = errors.New("missing")

func main() {
	err := fmt.Errorf("open settings: %w", ErrMissing)
	if errors.Is(err, ErrMissing) {
		fmt.Println("use defaults")
		return
	}
	fmt.Println("stop")
}
