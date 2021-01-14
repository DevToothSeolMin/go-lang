package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"ramen", "pizza"}
	sul := person{name: "sul", age: 32, favFood: favFood}
	fmt.Println(sul)
}
