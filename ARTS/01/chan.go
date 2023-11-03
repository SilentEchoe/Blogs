package main

import "os"

var (
	MaxWorker = os.Getenv("MAX_WORKERS")
	MaxQueue  = 1000
)

var Queue chan Payload

type Payload struct {
	// [redacted]
}

type Job struct {
	Payload
}

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func init() {
	Queue = make(chan Payload, MaxQueue)
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}
