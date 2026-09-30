package main

import (
	"fmt"
	"net/http"
)

func main() {
	request, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1:1", nil)
	_, err := http.DefaultClient.Do(request)
	fmt.Println(err != nil)
}
