# Concurrent Systems Programming with Go — lesson m05l04 — Parsing JSON with encoding/json
# https://learnsome.tech/courses/go-course/watch?lesson=m05l04
# © LearnSome.tech
package main

import (
	"encoding/json"
	"fmt"
)

type Report struct { Host string `json:"host"`; Healthy bool `json:"healthy"` }

func main() {
	report, err := json.Marshal(Report{Host: "api", Healthy: true})
	fmt.Println(string(report), err)
}
