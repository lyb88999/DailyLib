package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello world!666")
	if err != nil {
		return
	}
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	server := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	log.Fatalln(server.ListenAndServe())
}
