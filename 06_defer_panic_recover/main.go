package main

import "fmt"

// https://blog.golang.org/defer-panic-and-recover
func main() {
	// This function will be called when current function returns
	// Defer is used to simplify cleanup operations
	// It is guaranteed for this line to be executed, it is like "finally" keyword in java
	defer fmt.Println("world")

	// This statements will be added in stack
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("hello")
}
