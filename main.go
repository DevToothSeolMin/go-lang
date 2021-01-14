package main

import (
	"fmt"
)

func canDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}

	return false
}

func main() {
	fmt.Println(canDrink(18))
}
