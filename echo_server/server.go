package main

import (
	"fmt"
	"log"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/echo" { // Testing out path validation
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {  // Testing out r.Method
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Request contents will be echoed here")
}

func main() {
	http.HandleFunc("/echo", echoHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}