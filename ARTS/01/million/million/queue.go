package million

import (
	"fmt"
	"time"
)

var Queue chan Payload

type Job struct {
	Payload Payload
}

// JobQueue 发送工作请求的缓冲通道
var JobQueue chan Job

// Worker 工作池
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

var (
	MAXQueue  = 1000
	MaxWorker = 1000
)

func init() {
	Queue = make(chan Payload, MAXQueue)
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
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				if err := job.Payload.UploadToS3(); err != nil {
					fmt.Errorf("Error uploading to S3: %s", err.Error())
				}

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

// UploadToS3 上传到s3 存粗
func (p Payload) UploadToS3() error {
	// TODO 做一些上传的操作，这里为了方便使用线程休眠，模拟任务执行
	time.Sleep(3 * time.Second)
	return nil
}
