package main

import "fmt"

func main() {
	names := []string{"nico", "lynn", "dal"} //[]에 숫자 적으면 길이가 정해진 array, 숫자 안적으면 가변 배열 slice
	// names[3] = "alal" // 정해진 곳에 값 넣을 때
	// names[4] = "ie"
	names = append(names, "ala") // slice 현재 길이 이상에 값 추가할 때
	names[2] = "jeongsoep"
	fmt.Println(names)
}
