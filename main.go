package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))

	// log.Printf("Serving %s on HTTP port: %s\n", ".", "8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
