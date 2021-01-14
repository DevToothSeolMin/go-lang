package main

import (
	"fmt"
	"go-lang/banking"
)

func main() {
	account := banking.Account{Owner: "sul", Balance: 1000}
	fmt.Println(account)
}
