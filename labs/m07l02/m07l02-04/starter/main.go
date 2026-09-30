package main

import "fmt"

type Result struct { Name string; Healthy bool }

func main() {
	results := []Result{{"api", true}, {"web", false}}
	for _, result := range results {
		fmt.Println(result.Name, result.Healthy)
	}
}
