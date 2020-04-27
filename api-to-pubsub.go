package main

import (
	"fmt"
	"log"
	"net/http"
    "os"
)

const friendlyPackageName string = "API to Pub/Sub"

func fetchAndForward() {
    log.Print("called")
}

func handler(w http.ResponseWriter, r *http.Request) {
    fetchAndForward()
    fmt.Fprintf(w, "Done")
}

func main() {
	log.Printf("%s started.", friendlyPackageName)

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}