package main

import (
	"fmt"
)

func main() {

	x := 10
	y := 2

	// 算術演算子
	fmt.Println(x + y) // 加し算
	fmt.Println(x - y) // 引き算
	fmt.Println(x * y) // 掛け算
	fmt.Println(x / y) // 割り算
	fmt.Println(x % y) // 剰余算

	// 関係演算子
	fmt.Println(x > y) 
	fmt.Println(x < y)
	fmt.Println(x >= y) 
	fmt.Println(x <= y)
	fmt.Println(x == y) 
	fmt.Println(x != y)

	a := 8
	b := 3

	// 論理演算子
	fmt.Println(a >= 5 && a <= 10)
	fmt.Println(b >= 5 && b <= 10)
	fmt.Println(a >= 5 || b <= 10)
	fmt.Println(a >= 5 || b <= 10)

	// 代入演算子・複合代入演算子
	a += 10
	b += x
	fmt.Println(a)
	fmt.Println(b)

	a-- // インクリメント
	b-- // デクリメント

	fmt.Println(a)
	fmt.Println(b)

}