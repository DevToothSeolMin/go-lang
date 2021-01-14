package main

import (
	"fmt"
)

func main() {
	names := map[string]string{"name": "sul", "age": "32"}

	for key, value := range names {
		fmt.Println(key, value)
	}

	fmt.Println(names)
}
