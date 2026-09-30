package main

import (
	"fmt"
	"sync"
)

func main() {
	results := make(chan string, 2)
	var wait sync.WaitGroup
	for _, name := range []string{"api", "web"} {
		wait.Add(1)
		go func(target string) {
			defer wait.Done()
			results <- target + " checked"
		}(name)
	}
	wait.Wait(); close(results)
	for result := range results { fmt.Println(result) }
}
