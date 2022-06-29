package main

import (
	"log"
	"net/http"
)

func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
