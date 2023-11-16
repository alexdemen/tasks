package main

import "fmt"

func main() {
	var ch = make(chan string, 5)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- fmt.Sprintf("from %d", i)
		}
	}()

	for val := range ch {
		fmt.Println(val)
	}
}
