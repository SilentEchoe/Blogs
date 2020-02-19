// Gabriel Aszalos 协助完成这个示例
// runner 包管理处理任务的运行和生命周期
package main

import (
	"errors"
	"os"
	"os/signal"
	"time"
	)

// Runner 在给定的超时时间内执行一组任务
// 并且在操作系统发送中断信号时结束这些任务
type  Runner struct {
	// interrupt 通道报告从操作系统
	// 发送的信息
	interrupt chan  os.Signal

	// complete chan error
	complete chan  error

	//timeout 报告处理任务已经超时
	timeout <- chan time.Time

	//tasks 持有一组以索引顺序依次执行的函数
	tasks []func(int)
}


// ErrTimeout 会在任务执行超时时返回
var ErrTimeout = errors.New("received timout")

// ErrInterrupt 会在接收到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

// New 返回一个新的准备使用的Runner
func New(d time.Duration) *Runner  {
	return &Runner{
		interrupt: make(chan os.Signal,1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add将一个任务附加到Runner上。这个任务是一个
// 接受一个int类型的ID未作参数的函数
func (r *Runner) Add(tasks ...func(int))  {
	r.tasks = append(r.tasks,tasks...)
}

//Start 执行所有任务，并监视通道事件
func (r *Runner)Start() error  {
	// 我门希望接受所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同的goroutine 执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 当任务处理完成时发出的信号
	case err := <-r.complete:
		return err

	case <-r.timeout:
		return  ErrTimeout
	}


}

func (r *Runner) run() error  {
	for id, task := range r.tasks{
		// 检查操作系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// 执行已注册的任务
		task(id)
	}
	return  nil
}

func (r *Runner) gotInterrupt() bool{
	select {
	//当中断事件被触发时发出的信号
	case <- r.interrupt:
		// 停止接受后续的任何信号
		signal.Stop(r.interrupt)
		return  true
	// 继续正常运行
	default:
		return  false
	}

}

