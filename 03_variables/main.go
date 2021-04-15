package main

import (
	"fmt"
)

// Declaring of package-level variable with boolean type
// Default value is false
var global bool

// This is a block of different variables
var (
	// Unsigned int
	z   uint = 150
	zzz bool = true
)

func main() {
	// Declaring function-scope variables of type int
	// Default value is 0
	var i, j, k int

	fmt.Println(i, j, k, global)

	// This variable is initialized
	var a, b int = 1, 2
	fmt.Println(a, b)

	// Variable type is infered from artument
	var num, str, bin = 1, "Go", true
	fmt.Println(num, str, bin)

	// This is short declaration (without var keyword)
	// This construct is only available inside the method
	kk := 3
	fmt.Println(kk)

	// Float variable
	var f float32 = 1.5
	fmt.Println(f)
	fmt.Println(z, zzz)

	// Converting between types
	var negative int = -10
	// In this case we just drop a first sign byte, so we get 18446744073709551606
	z = uint(negative)
	fmt.Println(z)
	// z = negative - this is not allowed to assign different types, conversion is required

	// This is constant
	const world = "World!"
	fmt.Println(world)
}
