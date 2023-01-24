package main

import (
	"net/http"

	"github.com/akrylysov/algnhsa"
	"github.com/go-lambda/api/respond"
)

func main() {
	notFoundHandler := func(w http.ResponseWriter, r *http.Request) {
		respond.WithError(w, "not found", http.StatusNotFound)
	}
	algnhsa.ListenAndServe(http.HandlerFunc(notFoundHandler), nil)
}
