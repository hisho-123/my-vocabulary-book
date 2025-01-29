package main

import (
	router "backend/src/infra"
	"log"
)

func main() {
	r := router.NewRouter()
	
	log.Println("Server running on localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
