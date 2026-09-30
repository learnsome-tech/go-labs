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
