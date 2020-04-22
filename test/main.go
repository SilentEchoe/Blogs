package main

import "github.com/micro/go-micro"

func main() {
	service := micro.NewService(
		micro.Name("helloworld"))

	service.Run()
}
