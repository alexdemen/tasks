package main

import (
	"fmt"
)

func main() {
	var ch = make(chan struct{})
	go func() {
		fmt.Printf("Hello")
		ch <- struct{}{}
	}()
	<-ch
}
