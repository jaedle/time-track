package main

import (
	"fmt"
	"log"
	"net/http"
)

func HealthEndpoint(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/api/health", HealthEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
