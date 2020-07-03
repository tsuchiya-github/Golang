package main

import (
	"fmt"
)

func main() {
	// Golangにはwhile文がなく，for文のみ
	for i := 0; i <= 4; i++ {
		fmt.Println(i)
	}

	// break文
	for i := 0; i <= 4; i++ {
		if i == 3 {
			break // 3のときfor文を抜ける
		}
		fmt.Println(i)
	}

	// continue文
	for i := 0; i <= 4; i++ {
		if i == 3 {
			continue // 3のとき残りの処理をスキップ
		}
		fmt.Println(i)
	}

	// ネスト
	for i := 0; i <= 2; i++ {
		for j := 0; j <=2; j++ {
			fmt.Println(i, "-", j)
		}
	}

	// 配列操作
	array := [...]int{2, 4, 6, 8, 10}
	sum := 0

	for i := 0; i <=4; i++ {
		sum += array[i]
	}
	fmt.Println(sum)

	// for文はスタート値，継続条件，増減式が全て無くても書ける
	i := 0
	for i <= 4 {
		fmt.Println(i)
		i++
	}
	
}