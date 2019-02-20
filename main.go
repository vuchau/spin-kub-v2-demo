package main

import (
	"fmt"
	"net/http"
)

func Sum(x int, y int) int {
	return x + y

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling %+v\n", r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(" Hello world v1.2"))
}
func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	port := ":8000"
	fmt.Printf("Starting to service on port %s\n", port)
	http.ListenAndServe(port, nil)
}
