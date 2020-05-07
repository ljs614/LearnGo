package main

import (
	"fmt"

	"github.com/ljs614/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("jeongseop")
	account.Deposit(10000)
	fmt.Println(account)

}
