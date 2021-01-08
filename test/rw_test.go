package main

import (
	"sync"
	"time"
)

type RW interface {
	Write()
	Read()
}

const cost = time.Microsecond

type Lock struct {
	count int
	mu    sync.Mutex
}

func (l *Lock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost)
	l.mu.Unlock()
}

func (l *Lock) Read() {
	l.mu.Lock()
	time.Sleep(cost)
	_ = l.count
	l.mu.Unlock()
}
