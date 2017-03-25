package main

import (
	"net/http"
	"io"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Go!")
	return
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Starting server on port 3001")
	http.ListenAndServe(":3001", nil)
	return
}
