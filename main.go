package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling %+v\n", r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(" Hello world! updated "))
}

func main() {
	http.HandleFunc("/", index)
	port := ":8000"
	fmt.Printf("Starting to service on port %s\n", port)
	http.ListenAndServe(port, nil)
}
