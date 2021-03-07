package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func main() {

	fmt.Println("Server starting...")

	http.HandleFunc("/", hello)

	fmt.Println("Server started...")

	http.ListenAndServe(":1234", nil)
}
