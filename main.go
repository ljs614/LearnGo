package main

import "fmt"

func main() {
	js := map[string]string{"name": "jeongseop", "age": "30"} //map[key]value
	fmt.Println(js)
	for key, value := range js {
		fmt.Println(key, value)
	}
}
