package main

import (
	"fmt"
	"sync"
)

type treadSafety struct {
	mtx   sync.Mutex
	count int
}

func (t *treadSafety) change(count int) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	t.count = count
}

func main() {
	ts := treadSafety{mtx: sync.Mutex{}}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			ts.change(i)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(ts)
}
