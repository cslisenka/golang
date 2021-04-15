package main

import (
	"fmt"

	"rsc.io/quote"
)

/**
  go.mod - module file, when we create module
*/
func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println("End")
}

func sample() string {
	return "sample"
}
