package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf(
		"Goroutines: %d",
		runtime.NumGoroutine(),
	)
}
