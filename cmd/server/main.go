package main

import (
	"log"
	"net/http"

	"github.com/valeryvash/metrics/internal/handler"
	"github.com/valeryvash/metrics/internal/repository"
	"github.com/valeryvash/metrics/internal/service"
)

func main() {
	storage := repository.NewMemStorage()
	service := service.NewService(storage)
	handler := handler.NewHandler(service)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
