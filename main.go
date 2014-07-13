package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("KUDOS_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Kudos running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
