package main

import (
	"fmt"
	"sync"
)

func main() {
	group := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		group.Add(1)
		go func() {
			fmt.Println("id = ", i)
			group.Done()
		}()
	}

	group.Wait()
}
