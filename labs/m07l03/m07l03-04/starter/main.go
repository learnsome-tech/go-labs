package main

import "fmt"

type Result struct { Name string; Err error }

func main() {
	result := Result{Name: "api", Err: fmt.Errorf("timeout")}
	fmt.Println(result.Name, result.Err)
}
