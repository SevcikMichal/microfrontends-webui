package main

import (
	"log"
	"net/http"

	"github.com/SevcikMichal/microfrontends-webui/configuration"
	"github.com/SevcikMichal/microfrontends-webui/router"
)

func main() {
	log.Println("Main function starting...")
	startHTTPServer()
}

func startHTTPServer() {
	router := router.CreateRouter()

	server := &http.Server{
		Addr:    ":" + configuration.GetHttpPort(),
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err, "problem running server")
	}
}
