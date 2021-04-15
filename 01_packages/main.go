package main

import (
	"fmt"
	"math"
	"math/rand"

	"example.com/hello/01_packages/subpackage"
)

func main() {
	fmt.Println("M favorite number is", rand.Intn(10))
	// Only names with Capital letter are visible outside of the package
	// Names with no capital letter are not visible (protected)
	fmt.Println(math.Pi)

	// Use structure declared in the other package
	// Structure has only one public field, we only may access it
	it := subpackage.Item{Price: 1.1}
	subpackage.PublicFunc(it)
	// Try to print it
	fmt.Println(it)

	// We may not access private function

	// Package versioning
	// There is no standard support, we need to use package-management tool
	// TODO explore tools
}
