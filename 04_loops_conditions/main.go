package main

import (
	"fmt"
)

func main() {
	// General for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// For loop with just a condition
	sum := 0
	for sum < 1000 {
		sum += 100
	}
	fmt.Println(sum)

	var counter int = 0
	// Infinite loop
	for {
		fmt.Println(counter)
		counter++
		if counter > 3 {
			break
		}
	}

	// If with additional statement that is executed before a condition
	if v := 15; v > 10 {
		fmt.Println("15 > 10")
	}
}
