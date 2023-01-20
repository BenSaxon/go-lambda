package main

import (
	"github.com/akrylysov/algnhsa"
	"github.com/go-lambda/api/handlers/example/example"
	"github.com/joerdav/zapray"
)

func main() {
	log, err := zapray.NewProduction()
	if err != nil {
		panic("failed to create logger: " + err.Error())
	}

	handler, err := example.NewHandler(log)
	if err != nil {
		panic("failed to create handler: " + err.Error())
	}

	log.Info("handler created, listening")
	algnhsa.ListenAndServe(handler, nil)
}
