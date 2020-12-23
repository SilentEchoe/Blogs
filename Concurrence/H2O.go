package main

import (
	"context"
	"github.com/marusama/cyclicbarrier"
	"golang.org/x/sync/semaphore"
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"
)

// 定义水分子合成的辅助数据结构
type H2O struct {
	semaH *semaphore.Weighted         // 氢原子的信号量
	semaO *semaphore.Weighted         // 氧原子的信号量
	b     cyclicbarrier.CyclicBarrier // 循环栅栏，用来控制合成
}

func New() *H2O {
	return &H2O{
		semaH: semaphore.NewWeighted(2), //氢原子需要两个
		semaO: semaphore.NewWeighted(1), // 氧原子需要一个
		b:     cyclicbarrier.New(3),     // 需要三个原子才能合成
	}
}

func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	h2o.semaO.Acquire(context.Background(), 1)

	releaseHydrogen() // 输出H
	h2o.b.Await(context.Background())
	h2o.semaH.Release(1)
}

func (h2o *H2O) oxygen(releaseOxygen func()) {
	h2o.semaO.Acquire(context.Background(), 1)

	releaseOxygen()
	h2o.b.Await(context.Background())
	h2o.semaO.Release(1)
}

func TestWaterFactory(t *testing.T) {
	var ch chan string
	releaseHydrogen := func() { ch <- "H" }
	releaseOxygen := func() { ch <- "O" }
	var N = 100
	ch = make(chan string, N*3)

	h20 := New()

	var wg sync.WaitGroup
	wg.Add(N * 3)

	for i := 0; i < 2*N; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
			h20.hydrogen(releaseHydrogen)
			wg.Done()
		}()
	}

	for i := 0; i < 2*N; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
			h20.oxygen(releaseOxygen)
			wg.Done()
		}()
	}

	wg.Wait()

	if len(ch) != N*3 {
		t.Fatalf("expect %d atom but got %d", N*3, len(ch))
	}

	var s = make([]string, 3)
	for i := 0; i < N; i++ {
		s[0] = <-ch
		s[1] = <-ch
		s[2] = <-ch
		sort.Strings(s)

		water := s[0] + s[1] + s[2]
		if water != "HHO" {
			t.Fatalf("expect a water molecule but got %s", water)
		}
	}

}
