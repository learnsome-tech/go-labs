# Concurrent Systems Programming with Go — lesson m07l04 — Exporting Results as JSON
# https://learnsome.tech/courses/go-course/watch?lesson=m07l04
# © LearnSome.tech
package main

import (
	"encoding/json"
	"fmt"
)

type Result struct { Name string `json:"name"`; Healthy bool `json:"healthy"` }

func main() {
	data, _ := json.Marshal([]Result{{"api", true}})
	fmt.Println(string(data))
}
