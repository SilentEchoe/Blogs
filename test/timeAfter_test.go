package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func doBadthing(done chan bool) {
	time.Sleep(time.Second)
	done <- true
}

func timeout(f func(chan bool)) error {
	done := make(chan bool)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Microsecond):
		return fmt.Errorf("timeout")
	}

}

func test(t *testing.T, f func(chan bool)) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		timeout(f)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}

func TestBadTimeout(t *testing.T) {
	test(t, doBadthing)
}
