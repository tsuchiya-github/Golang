package main

import (
	"fmt"
)

func main() {
	
	// 一次元配列
	array01 := [3]string{"sato", "suzuki", "takahashi"}
	// array01 := [...]string{"sato", "suzuki", "takahashi"} // [...]で要素数省略可能

	fmt.Println(array01[0])
	fmt.Println(array01[1])
	fmt.Println(array01[2])

	array01[1] = "tsuchiya"
	fmt.Println(array01[0])
	fmt.Println(array01[1])
	fmt.Println(array01[2])

	// 二次元配列
	array02 := [2][2]string{{"sato", "suzuki"}, {"takahashi", "tsuchiya"}}
	fmt.Println(array02[0][0])
	fmt.Println(array02[0][1])
	fmt.Println(array02[1][0])
	fmt.Println(array02[1][1])

}