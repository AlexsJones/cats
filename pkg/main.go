package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	started = time.Now()
)

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func main() {

	fileServer := http.FileServer(http.Dir("./static")) // New code
	http.Handle("/", fileServer)
	http.Handle("/healthz", http.HandlerFunc(health))

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
