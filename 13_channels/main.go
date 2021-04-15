package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(c chan string) {
	// Generates strings and sends to channel
	for i := 0; i < 100; i++ {
		message := fmt.Sprintf("message-%d", rand.Int())
		// Try sending, if not possible (buffer is full) - drop (kind of a backpressure)
		// If we do not use select we block here until consumer gets message
		// We can use select with multiple channels, select blocks until the first chan becomes available
		// The first availavle channel is chosen in this case
		select {
		// Send message to channel
		case c <- message:
			fmt.Println("Produced", message)
			fmt.Println("Messages in channel", len(c))
		case t := <-time.After(time.Millisecond * 300):
			// Wait for 1 second, if we blocked for more time - timeout
			fmt.Println("Times out", t)
			// If we use default we drop immediately
			//default:
			//	fmt.Println("Dropped")
		}
	}
}

func consumer(c chan string) {
	// Receives data and displays
	for {
		// Receive message from channel
		message := <-c
		fmt.Println("Received ", message)
		// Sleep for 1 second
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	fmt.Println("Channels demo")

	// Declare channel with buffer of 100 messages
	test := make(chan string, 10)

	// Start consumer
	go consumer(test)

	// Start producer
	producer(test)
}
