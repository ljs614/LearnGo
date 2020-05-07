package main

import (
	"fmt"

	"github.com/ljs614/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("jeongseop")
	fmt.Println(account)
}
