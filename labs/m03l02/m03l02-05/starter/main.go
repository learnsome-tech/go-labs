package main

import (
	"errors"
	"fmt"
)

var ErrBusy = errors.New("busy")

func start() error {
	return fmt.Errorf("start worker: %w", ErrBusy)
}

func main() {
	err := start()
	fmt.Println(err)
	if errors.Is(err, ErrBusy) {
		fmt.Println("retry later")
	}
}
