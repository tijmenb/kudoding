package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("KUDOS_PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Kudos running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
