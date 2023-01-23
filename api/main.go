package main

import (
	"net/http"

	"github.com/go-lambda/api/handlers/example"
	"github.com/gorilla/mux"
	"github.com/joerdav/zapray"
)

func main() {
	log, err := zapray.NewDevelopment()
	if err != nil {
		panic("failed to create logger: " + err.Error())
	}

	log.Info("creating handlers")
	exampleHandler, err := example.NewHandler(log)
	if err != nil {
		panic("failed to get example handler: " + err.Error())
	}

	log.Info("initializing router")
	router := mux.NewRouter()
	router.Handle("/example/{message}", exampleHandler)

	log.Info("Starting server")
	panic(http.ListenAndServe("127.0.0.1:8080", router))
}
