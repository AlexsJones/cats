package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	started        = time.Now()
	Version string = "dev"
)

func health(w http.ResponseWriter, r *http.Request) {
	duration := time.Now().Sub(started)
	if duration.Seconds() > 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}

func main() {

	fileServer := http.FileServer(http.Dir("./static")) // New code
	http.Handle("/", fileServer)
	http.Handle("/healthz", http.HandlerFunc(health))

	fmt.Printf("Starting server at port 8080\n")
	fmt.Printf("Version %s\n", Version)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
