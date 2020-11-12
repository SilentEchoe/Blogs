package main

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

func publish(tn string, msg interface{}) error {
	var newTask = Task{
		Topic:   tn,
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
		}
	}
	if isExist {
		MySubscription = append(MySubscription, topic)
	}

	msgChan := make(chan interface{})
	for _, value := range MyTaskQueue {
		if value.Topic == topic {
			msgChan <- value
		}
	}
	return msgChan, nil
}

func unsubscribe(topic string) error {
	for _, value := range MySubscription {
		if value == topic {

		}
	}
	var index = 0
	for i := 0; i <= len(MySubscription)-1; i++ {
		if MySubscription[i] == topic {
			index = i
			break
		}
	}
	MySubscription = append(MySubscription[:index], MySubscription[index+1:]...)
	return nil
}
