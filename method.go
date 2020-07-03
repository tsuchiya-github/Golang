package main

import (
	"fmt"
)

type Student struct{
	name          string
	math, english float64
}

func (s Student) avg1() { //引数のないmethod
	fmt.Println(s.name, (s.math + s.english) / 2)
}

func (s Student) avg2(math, english float64){ //引数のあるmethod
	fmt.Println(s.name, (s.math + s.english) / 2)
}

func (s Student) avg3(math, english float64) float64{ // 戻り値のあるmethod
	return (math + english) / 2
}

func (s Student) avg4(math, english float64) (avgResult float64) { // returnの変数を省略するmethod
	avgResult = (math + english) / 2
	return
}

func main() {
	// メソッドとは構造体などの特定の方に関連づけられた関数のこと
	a001 := Student{"sato", 80, 70}
	a001.avg1()
	a001.avg2(80, 70)
	fmt.Println(a001.avg3(80, 70))
	fmt.Println(a001.avg4(80, 70))

}