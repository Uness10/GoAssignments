package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func increment() {
	mu.Lock()   // Acquire the lock
	counter++   // Critical section
	mu.Unlock() // Release the lock
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Final Counter Value: %d\n", counter)
}
