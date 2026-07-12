package main

import (
	"log"

	"github.com/rajabhishekmaurya/ecommerce-microservices/payment-service/internal/server"
)

func main() {

	srv := server.New()

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
