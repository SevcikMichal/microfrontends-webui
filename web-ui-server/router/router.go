package router

import (
	"github.com/SevcikMichal/microfrontends-webui/server"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// basePathRouter := router.PathPrefix(configuration.GetBaseURL()).Subrouter()

	router.HandleFunc("/", server.ServeSinglePageApplication).Methods("GET")

	return router
}
