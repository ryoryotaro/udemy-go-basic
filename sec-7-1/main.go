package main

import "fmt"

//関数

// x + y の結果が int 型で、これが戻り値
func Plus(x int, y int) int {
	//func Plus(x , y int) int { と省略可能
	return x + y

}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) { //戻り値の変数名を最初に定義
	result = price * 2
	// 戻り値を省略できる
	return
}

func Noreturn() {
	fmt.Println("No Return")
	return
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(9, 4)
	fmt.Println(i2, i3)

	i4, _ := Div(9, 4)
	fmt.Println(i4)

	i5 := Double(1000)
	fmt.Println(i5)

	Noreturn()

}
