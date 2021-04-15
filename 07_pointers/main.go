package main

import (
	"fmt"
)

func increment(val *int) {
	// Increment value
	// Get value by pointer
	*val++
}

func main() {
	var a int = 0
	var pointer *int = &a
	increment(pointer)

	fmt.Println(a)
	// Print pointer itself
	fmt.Println(pointer)
	// Print value by pointer
	fmt.Println(*pointer)
}
