package router

import (
	"github.com/SevcikMichal/microfrontends-webui/server"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// basePathRouter := router.PathPrefix(configuration.GetBaseURL()).Subrouter()

	// index.html
	router.HandleFunc("/", server.ServeSinglePageApplication).Methods("GET")

	// modules, fonts, assets
	router.PathPrefix("/modules").HandlerFunc(server.ServeFile).Methods("GET")
	router.PathPrefix("/assets").HandlerFunc(server.ServeFile).Methods("GET")
	router.PathPrefix("/fonts").HandlerFunc(server.ServeFile).Methods("GET")

	return router
}
