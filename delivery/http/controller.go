package http

import (
	"github.com/gorilla/mux"
	"github.com/hobord/golang-poc-rest/delivery/http/handlers"
	"github.com/hobord/golang-poc-rest/usecase"
)

// MakeRouting is add handler functions to mux router
func MakeRouting(router *mux.Router, fooInteractor usecase.FooInteractorInterface) {
	fooApp := handlers.CreateFooRestHTTPModule(fooInteractor)

	router.HandleFunc("/foo", fooApp.Create).Methods("POST")
	router.HandleFunc("/foo/{id}", fooApp.GetByID)
	router.HandleFunc("/foo", fooApp.GetAll).Methods("GET")
	router.HandleFunc("/foo", fooApp.Update).Methods("PUT")
	router.HandleFunc("/foo/{id}", fooApp.Delete).Methods("DELETE")
}
