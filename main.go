package main

import (
	"fmt"
)

func canDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}

	return true
}

func main() {
	fmt.Println(canDrink(16))
}
