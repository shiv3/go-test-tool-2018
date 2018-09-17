package main

import (
	"github.com/shiv3/go-integration-test/handle"
	"net/http"
)

func main() {
	h := handle.Handler()
	http.ListenAndServe(":8080", h)
}
