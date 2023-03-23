package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting: Building Apps for the k8 app")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received from %s", r.RemoteAddr)
	fmt.Fprintf(w, "Building Apps for k8 app says Hi")
}
