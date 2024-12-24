package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", rolldice)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func rolldice(w http.ResponseWriter, r *http.Request) {
	roll := 1 + rand.Intn(6)

	resp := strconv.Itoa(roll) + "\n"
	if _, err := io.WriteString(w, resp); err != nil {
		log.Printf("Write failed: %v\n", err)
	}
}
