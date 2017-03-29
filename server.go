package main

import (
	"net/http"
	"io"
	"fmt"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Go!")
	return
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Starting server on port 3001")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	http.ListenAndServe(":"+port, nil)
	return
}
