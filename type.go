// Golangでは変数にデータを入れる際にデータ型を指定する必要がある(静的型付け言語という:Java,Kotlin,Cなど)

package main

import (
	"fmt"
	"reflect" // reflectパッケージを用いてデータ型を確認する
)

func main() {
	
	// 数値型
	num01 := 123 // データ型省略
	var num02 int = 1234567890 // int型宣言
	num03 := 1.23
	var num04 float64 = 1.23456789 // float64型宣言

	fmt.Println(num01)
	fmt.Println(reflect.TypeOf(num01))

	fmt.Println(num02)
	fmt.Println(reflect.TypeOf(num02))

	fmt.Println(num03)
	fmt.Println(reflect.TypeOf(num03))

	fmt.Println(num04)
	fmt.Println(reflect.TypeOf(num04))

	// 文字列型
	var string_a string = "Hello, World!" // string型宣言
	string_b := "Hello, World!" // データ型省略
	
	fmt.Println(string_a)
	fmt.Println(reflect.TypeOf(string_a))

	fmt.Println(string_b)
	fmt.Println(reflect.TypeOf(string_b))

	// ブール型
	a := 10
	b := 1
	num_bool := a < b // ブール型宣言
	
	fmt.Println(num_bool)
	fmt.Println(reflect.TypeOf(num_bool))

}