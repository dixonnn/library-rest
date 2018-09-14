package main

import (
	"log"
	"net/http"
)

// Main function starts up Router and listens on :8080
func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
