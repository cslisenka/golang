package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on")

	os := runtime.GOOS
	switch os {
	case "linux":
		fmt.Println("Linux!")
	default:
		fmt.Println(os)
	}

	// Switch with no conditions, helps to avoid ugly if-else-else-if statements
	switch {
	case 5 < 10:
		fmt.Println("5 < 10")
	default:
		fmt.Println("Default")
	}
}
