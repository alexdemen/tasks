package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	ticker := time.NewTicker(3 * time.Second)

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-ticker.C:
			fmt.Println("tick")
		default:
			time.Sleep(time.Second * 1)
			ch <- 42
		}
	}

}
