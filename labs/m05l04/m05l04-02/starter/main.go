package main

import (
	"encoding/json"
	"fmt"
)

type Reply struct { Status string `json:"status"`; Count int `json:"count"` }

func main() {
	var reply Reply
	err := json.Unmarshal([]byte(`{"status":"ok","count":2}`), &reply)
	fmt.Println(reply.Status, reply.Count, err)
}
