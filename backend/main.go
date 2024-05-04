package main

import (
	"fmt"
	"log"
	"net/http"
)

const ServerPort = ":8080"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /boo", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "HELLO\n")
  })

	mux.HandleFunc("GET /boo/{id}", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    fmt.Fprintf(w, "handling task with id=%v\n", id)
  })

	fmt.Println("The server is on tap at http://localhost:8080.")
	log.Fatal(http.ListenAndServe(ServerPort, mux))
}