package main

import (
	"fmt"
)

// 定数
// 他のパッケージからも呼び出したい場合は頭文字を大文字にする
const Pi = 3.14

const (
	URL      = "http://example.com"
	SiteName = "test"
)

// 値を省略することができる
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

var Big int = 9222222223232322 + 1

const (
	c0 = iota //整数の連番を作る
	c1
	c2
)

func main() {
	fmt.Println(Pi)
	fmt.Println(URL, SiteName)
	fmt.Println(A, B, C, D, F, F)
	fmt.Println(Big - 1)
	fmt.Println(c0, c1, c2)

}
