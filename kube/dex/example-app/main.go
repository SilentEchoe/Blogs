package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/login", DexLogin)

	// Start the server on port
	http.ListenAndServe(":5555", nil)
}

func DexLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
