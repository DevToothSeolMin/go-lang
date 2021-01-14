package main

import (
	"fmt"
)

func canDrink(age int) bool {
	if age < 18 {
		return false
	}

	return true
}

func main() {
	fmt.Println(canDrink(16))
}
