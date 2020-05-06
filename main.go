package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "chicken"}
	js := person{name: "jeongseop", age: 30, favFood: favFood}
	fmt.Println(js)
}
