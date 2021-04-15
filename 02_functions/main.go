package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

// Alternative way of declaring parameters with same type
func add2(x, y int) int {
	return x + y
}

// This function returns 2 results
// Function may return any number of results
func minMax(x, y int) (int, int) {
	if x > y {
		return y, x
	} else {
		return x, y
	}
}

// This function returns 3 results
// Function may return any number of results
func minMaxAvg(x, y int) (int, int, int) {
	avg := (x + y) / 2

	if x > y {
		return y, avg, x
	} else {
		return x, avg, y
	}
}

// This function returns named variales
func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min, max = y, x
	} else {
		min, max = x, y
	}

	// this is return statement without arguments (naked return)
	// we return min and max, so no need to mention it here
	return
}

func executor(fn func(int, int) int, a int, b int) {
	fmt.Println("Arguments: ", a, " ", b)
	fmt.Println("Function ", fn)
	result := fn(a, b)
	fmt.Println("result = ", result)
}

// This function accumulates
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func vararg(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func main() {
	fmt.Println(add(10, 15))
	fmt.Println(add2(10, 15))
	fmt.Println(minMax(10, 15))

	min, max := minMax(10, 15)
	fmt.Printf("Min = %d, Max = %d \n", min, max)

	fmt.Println(minMaxAvg(5, 10))
	fmt.Println(namedMinMax(5, 10))

	// Functions are values, we can pass then into the method
	sum := func(x, y int) int {
		return x + y
	}

	product := func(x, y int) int {
		return x * y
	}

	executor(sum, 2, 3)
	executor(product, 2, 3)

	// Functions may be closures
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	// Calling function with variable length of artuments
	vararg(1, 2, 3, 8, 9, 10)
}
