package main

import "fmt"

type Target struct { Name string; URL string }

func main() {
	targets := []Target{{"api", "https://api"}, {"web", "https://web"}}
	for _, target := range targets {
		fmt.Println(target.Name, target.URL)
	}
}
