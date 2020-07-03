package main

import (
	"fmt"
)

func sayHelloWorld() { // 引数なし
	fmt.Println("Hello, World!")
}

func sayHello(greeting string) { // 引数あり
	fmt.Println(greeting)
}
func calc(a int, b int) {
	fmt.Println(a * b)
} 

func calc_return(a int, b int) int { // 戻り値あり
	return (a / b)
}

func calc_return2(a int, b int) (int, int) { // 戻り値２つあり
	return (a / b), (a * b)
}

func calc_return3(a int, b int) (x int, y int) { // 関数内で変数を扱うことも可能
	x = a / b
	y = a * b
	return //戻り値を宣言しているため，return x, yと書かなくても大丈夫(書いてもいい)
}

func main() {
	sayHelloWorld()
	sayHello("Good morning!")
	calc(6, 3)
	fmt.Println(calc_return(6, 3))
	fmt.Println(calc_return2(6, 3))
	fmt.Println(calc_return3(6, 3))

	// main関数内で関数を定義することも可能
	hello := func(greeting string) { // funcのあとに関数名がない(無名関数)
		fmt.Println(greeting)
	}
	hello("Good morning!")

	//無名関数は以下のように書くことも可能
	func(greeting string) { // funcのあとに関数名がない(無名関数)
		fmt.Println(greeting)
	}("Good morning!")
}