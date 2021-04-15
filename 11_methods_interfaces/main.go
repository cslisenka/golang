package main

import (
	"fmt"
)

type Logger interface {
	// We can have more than 1 method here
	Log(message string)
}

// Empty struct, just to implement methods
type ConsoleLogger struct {
}

func (*ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

// Empty struct, just to implement methods
type BeautyLogger struct {
}

func (*BeautyLogger) Log(message string) {
	fmt.Println("<", message, ">")
}

func main() {
	// Declare variable of the interface type
	var cl Logger
	// Assign one implementation
	cl = &ConsoleLogger{}
	cl.Log("test")

	// Assign second implementation
	cl = &BeautyLogger{}
	cl.Log("test")
}
