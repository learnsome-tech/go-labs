# Concurrent Systems Programming with Go — lesson m07l04 — Exporting Results as JSON
# https://learnsome.tech/courses/go-course/watch?lesson=m07l04
# © LearnSome.tech
package main

import (
	"encoding/json"
	"os"
)

func main() {
	report := map[string]bool{"healthy": true}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(report)
}
