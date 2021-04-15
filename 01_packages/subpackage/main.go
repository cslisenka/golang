package subpackage

import (
	"fmt"
)

type Item struct {
	// This attribute is visible outside of the package
	Price float64
	// This attribute is not visible outside of the package
	count int
}

func PublicFunc(item Item) {
	fmt.Println(item.Price, item.count)
}

func private_func() {
	fmt.Println("private")
}
