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

func webhook(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "2137067855")
	if r.FormValue("hub.mode") == "subscribe" && r.FormValue("hub.verify_token") == "qwerty123" {
		io.WriteString(w, r.FormValue("hub.challenge"))
	} else {
		io.WriteString(w, "Verification Failed")
	}
	return
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/webhook", webhook)
	fmt.Println("Starting server on port 3001")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	http.ListenAndServe(":"+port, nil)
	return
}
