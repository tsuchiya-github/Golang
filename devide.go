package main

import (
	"fmt"
)

func main() {

	// if文
	age := 22

	if age >= 20 { // 20歳以上のとき
		fmt.Println("adult")
	} else if age == 0 { // 0歳のとき
		fmt.Println("baby")
	} else {
		fmt.Println("child") // それ以外のとき
	}

	// if文内で宣言も可能
	if age := 18 ; age >= 20 { // 20歳以上のとき
		fmt.Println("adult")
	} else if age == 0 { // 0歳のとき
		fmt.Println("baby")
	} else {
		fmt.Println("child") // それ以外のとき
	}
	
}