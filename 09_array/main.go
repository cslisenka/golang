package main

import (
	"fmt"
)

type User struct {
	name    string
	surname string
}

func main() {
	// Arrays are fixed-length sequences of items of the same type.
	// Assigning one array to another copies all of the elements
	// If you pass an array to a function, it will receive a copy of the array (not a pointer or reference to it)
	// Arrays are expensive to reassign!
	// Declare array on 10 elements
	var a [10]int
	a[0] = 1
	a[1] = 2

	fmt.Println(a)

	// Slice is a view of array, when we change array it is reflected in slice
	// Slices are reference types, it is cheap to pass it into function calls
	// A slice is a data structure describing a contiguous section of an array stored separately from the slice variable itself.
	// A slice is not an array. A slice describes a piece of an array.
	// Slise is like a data structure with 2 elements: poiner and length
	// https://blog.golang.org/slices
	// https://blog.golang.org/slices-intro
	var sl []int = a[1:3]
	fmt.Println(sl)

	a[1] = 3
	fmt.Println(sl)

	// We may edit slice, and changes are reflected in array
	sl[1] = 99
	fmt.Println(a, sl)

	// Init array when creating
	cities := [4]string{"Minsk", "Kiev", "Warsaw", "Vilnius"}
	fmt.Println(cities)

	// Create slice (do not specify length)
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}

	// Create slice (do not specify length)
	u := []User{
		{"aaa", "bbb"},
		{"ccc", "ddd"},
	}

	fmt.Println(s, u)

	// Use default end
	fmt.Println(u[1:])
	fmt.Println(u[:])

	// Print slice/array length
	println(len(a), " ", len(u), " ", len(cities[2:]))

	// Emoty slice
	var slice []int
	// Nil slice has length of 0
	fmt.Println(slice)
	if slice == nil {
		fmt.Println("slice is nill")
	}

	// Create dynamically sized array and return a slice that points to it
	sl3 := make([]int, 0, 10) // capacity 10
	fmt.Println("Slice with capacity of 10")
	fmt.Println(sl3)
	// Add new element to slise, here we produce new slice with double length
	sl3 = append(sl3, 1)
	// This method returns new slice (it may reconstruct the object)
	sl3 = append(sl3, 2)
	fmt.Println(sl3, " length ", len(sl3))

	// Slices of slices
	board := [][]int{
		{1, 2, 3},
		{1, 2},
		{1, 3, 3},
	}

	fmt.Println(board)

	// Add new element to slice
	newUser := User{"a", "b"}
	// Construct new slice by appending to existing one
	slice2 := append(u, newUser)
	fmt.Println(u, slice2)

	// Iterate over slice or array
	for i, v := range u {
		// We have i - iteration number, v - value
		fmt.Println(i, " ", v)
	}

	// We can skip i if we don't need
	for _, v := range u {
		// We have v - value
		fmt.Println(v)
	}
}
