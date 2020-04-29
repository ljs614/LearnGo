package main

import (
	"fmt"
	"strings"
)

// func lenAndUpper(name string) (int, string) { //... arg 여러개 받을 수 있음.
// 	return len(name), strings.ToUpper(name)
// }

func lenAndUpper(name string) (length int, uppercase string) { //naked return
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	len, upperName := lenAndUpper("jeongseop")
	onlyLen, _ := lenAndUpper("lee") //_ 나머지 무시
	fmt.Println(len, upperName)
	fmt.Println(onlyLen)

	repeatMe("nico", "lynn", "dal", "marl")
}
