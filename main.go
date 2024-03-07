package main

// @title Nat Service
// @version 1.0
// @description Description of the service

// @host localhost
// @schemes http

import (
	"github.com/abaykerimov/test_kmf/internal/app/handlers"
	"github.com/abaykerimov/test_kmf/internal/infrastructure/dependencies"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	di := dependencies.NewDI()
	handler := handlers.Handler{DI: di}

	router := mux.NewRouter()
	handler.LoadRoutes(router)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal("error")
	}
}
