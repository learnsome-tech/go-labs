package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	for _, name := range []string{"api", "worker"} {
		wait.Add(1)
		go func(target string) {
			defer wait.Done()
			fmt.Println(target)
		}(name)
	}
	wait.Wait()
}
