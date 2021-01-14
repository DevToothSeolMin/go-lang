package main

import (
	"fmt"
)

func main() {
	names := [5]string{"sul", "min", "del"}
	names[3] = "all"
	names[4] = "Dall"
	fmt.Println(names)
}
