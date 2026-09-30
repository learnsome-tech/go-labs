package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	health(recorder, request)
	fmt.Println(recorder.Code, recorder.Body.String())
}
