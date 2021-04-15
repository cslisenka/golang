package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter = 0
	// We can use read-write lock (lock.RWMutex)
	lock sync.Mutex
)

func process() {
	// We may lock on this mutex
	lock.Lock()
	// Guaranteed to be executed at the end of the method
	defer lock.Unlock()

	counter++
	fmt.Println("process!", counter)
}

func main() {
	for i := 0; i < 100; i++ {
		// Each function goes to the queue
		go process()
	}

	// This function call is last in the queue, so processed first
	// Inline function
	go func() {
		fmt.Println("test!")
	}()

	// TODO get end of all goroutines
	time.Sleep(time.Millisecond * 100)
	fmt.Println("counter = ", counter)
}
