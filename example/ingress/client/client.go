package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func hello(wr http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("service-test.default.svc.cluster.local:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}

	wr.Write([]byte("client succeed"))
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
