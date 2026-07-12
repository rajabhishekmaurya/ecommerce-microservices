package main

import (
	"log"
	"net/http"

	"github.com/rajabhishekmaurya/ecommerce-microservices/order-service/internal/server"
)

func main() {

	srv, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := srv.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
