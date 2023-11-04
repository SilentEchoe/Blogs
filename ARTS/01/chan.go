package main

import (
	"context"
	"google.golang.org/appengine/log"
)

var (
	MaxWorker = 1000
	MaxQueue  = 1000
)

var Queue chan Payload

type Payload struct {
}

type PayloadCollection struct {
	WindowsVersion string    `json:"version"`
	Token          string    `json:"token"`
	Payloads       []Payload `json:"data"`
}

type Job struct {
	Payload
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func init() {
	Queue = make(chan Payload, MaxQueue)
}

func (p *Payload) Work() error {
	// TODO just do work
	return nil
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				if err := job.Payload.Work(); err != nil {
					log.Errorf(context.Background(), "Error uploading to S3: %s", err.Error())
				}

			case <-w.quit:
				// signal to stop
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

// SimulationDo 模拟工作
func SimulationDo() {
	// 模拟一个请求
	var content = &PayloadCollection{}

	for _, payload := range content.Payloads {
		work := Job{Payload: payload}

		JobQueue <- work
	}
}

// 调度器

type Dispatcher struct {
	//向dispatcher注册的worker通道池
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool}
}

func (d *Dispatcher) Run() {
	for i := 0; i < cap(d.WorkerPool); i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	go d.dispatch()
}

// 调度
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}

func main() {
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
}
