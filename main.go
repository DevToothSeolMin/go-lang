package main

import (
	"fmt"
	"go-lang/accounts"
)

func main() {
	account := accounts.NewAccount("sul")
	fmt.Println(account)
}
