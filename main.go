package main

import (
	"fmt"
)

// func lenAndUpper(name string) (int, string) { //... arg 여러개 받을 수 있음.
// 	return len(name), strings.ToUpper(name)
// }

// func lenAndUpper(name string) (length int, uppercase string) { //naked return
// 	defer fmt.Println("I'm done")
// 	length = len(name)
// 	uppercase = strings.ToUpper(name)
// 	return
// }

// func repeatMe(words ...string) {
// 	fmt.Println(words)
// }

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers { //변수 하나만 하면 인덱스만 출력.
		total += number
	}
	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
