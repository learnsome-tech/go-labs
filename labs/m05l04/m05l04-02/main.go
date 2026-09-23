# Concurrent Systems Programming with Go — lesson m05l04 — Parsing JSON with encoding/json
# https://learnsome.tech/courses/go-course/watch?lesson=m05l04
# © LearnSome.tech
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
