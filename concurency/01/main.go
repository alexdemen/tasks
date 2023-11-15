package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mt := sync.Mutex{}
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go func(wg sync.WaitGroup) {
			mt.Lock()
			defer mt.Unlock()
			ch <- i
		}(wg)
	}

	for {
		select {
		case v := <-ch:
			fmt.Println("from goroutine ", v)
		}
	}
}
