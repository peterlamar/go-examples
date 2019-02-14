package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var name string

	found := r.URL.Query().Get("name")
	if found != "" {
		name = found
	} else {
		name = "world"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	fmt.Print("Go to http://localhost:3000/?name=Alice\n")
	fmt.Print("Or\n")
	fmt.Print("curl -H \"Content-Type: application/xml\" -X GET http://localhost:3000/?name=Alice \n")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}