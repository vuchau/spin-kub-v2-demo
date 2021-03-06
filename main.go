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
	w.Write([]byte(" Hello world, updated docker trigger 12347890 dev1!"))
}

func main() {
	http.HandleFunc("/", index)
	port := ":8000"
	fmt.Printf("Starting to service on port %s\n", port)
	http.ListenAndServe(port, nil)
}
