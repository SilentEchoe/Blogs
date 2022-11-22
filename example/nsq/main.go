package main

import (
	"fmt"

	nsq "github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()

	c, err := nsq.NewConsumer("nsq", "consumer", config)
	if err != nil {
		fmt.Println("Failed to init consumer: ", err.Error())
		return
	}

	c.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
		fmt.Println("received message: ", string(m.Body))
		m.Finish()
		return nil
	}))

	err = c.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		fmt.Println("Failed to connect to nsqd: ", err.Error())
		return
	}

	<-c.StopChan
}
