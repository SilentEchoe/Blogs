package main

import (
	"fmt"
	"sync"
)

type Task struct {
	Topic   string
	msg     interface{}
	IsExist int
}

type Subscription interface {
	publish(topic string, msg interface{}) error
	subscribe(topic string) (chan interface{}, error)
	unsubscribe(topic string) error
	close()
	broadcast(msg interface{}, subscribers []chan interface{})
}

var MyTaskQueue = make([]Task, 0)
var MySubscription = make([]string, 0)

func publish(topic string, msg interface{}) error {
	var newTask = Task{
		Topic:   topic,
		msg:     msg,
		IsExist: 0,
	}
	MyTaskQueue = append(MyTaskQueue, newTask)
	return nil
}

func subscribe(topic string) (chan interface{}, error) {
	var isExist = false

	for _, value := range MySubscription {
		if value == topic {
			isExist = false
			break
		} else {
			isExist = true
		}
	}

	if len(MySubscription) == 0 || isExist {
		MySubscription = append(MySubscription, topic)
	}

	msgChan := make(chan interface{}, len(MyTaskQueue))

	go func() {
		for _, value := range MyTaskQueue {
			if value.Topic == topic {
				msgChan <- value
			}
		}
		close(msgChan)
	}()

	return msgChan, nil
}

func unsubscribe(topic string) error {
	var index = 0
	for i := 0; i <= len(MySubscription)-1; i++ {
		if MySubscription[i] == topic {
			index = i
			break
		}
	}
	if index != 0 {
		MySubscription = append(MySubscription[:index], MySubscription[index+1:]...)
	}
	return nil
}

func broadcast(msg interface{}, subscribers []chan interface{}) {

}

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	_ = publish("Select", "show log")
	_ = publish("Delect", "Delect this log")
	go NewDoSelectTask("Select")
	go NewDoSelectTask("Delect")
	wg.Wait()
}

func NewDoSelectTask(topic string) {
	var consumer, err = subscribe(topic)
	if err == nil {
		defer wg.Done()
		select {
		case res := <-consumer:
			fmt.Println(res)
		}

	}
}
