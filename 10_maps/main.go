package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	var result map[string]int = make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		count := result[word]
		if count == 0 {
			result[word] = 1
		} else {
			result[word] = count + 1
		}

	}

	return result
}

func main() {
	var m map[string]int = make(map[string]int)

	m["Minsk"] = 1
	m["Vitebsk"] = 150

	fmt.Println(m)

	// Building a map using literals
	var m2 = map[string]int{
		"Vilnius": 1,
		"Kiev":    2,
	}
	fmt.Println(m2)

	// Delete element
	delete(m, "Minsk")

	// Test if element is presented
	elem, presented := m["N/a"]
	fmt.Println(elem, presented)

	fmt.Println(WordCount("apple apple sun moon house"))
}
