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
