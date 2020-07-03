package main

import (
	"fmt"
)

func main() {
	var s Student // 構造体の初期化
	s.name = "sato"
	s.math = 80
	s.english = 70
	fmt.Println(s)

	t := Student{"suzuki", 40, 60} // 構造化の初期化と同時にフィールドに値を代入
	fmt.Println(t)

	u := Student{name: "tsuchiya", math: 50, english: 90} // 構造化の初期化と同時にフィールドに値を代入(その2)
	fmt.Println(u)

}

// 構造体
type Student struct{
	name string
	math, english float64
}
