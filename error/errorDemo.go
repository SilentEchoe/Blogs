package main

import "log"

func main() {
	if r := recover(); r != nil {
		log.Fatal(r)
	}
	panic(123)

}
