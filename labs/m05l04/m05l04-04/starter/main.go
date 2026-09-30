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
