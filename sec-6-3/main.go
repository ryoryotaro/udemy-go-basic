package main

import "fmt"

//論理演算子

func main() {
	// 論理積
	fmt.Println(true && false == true)
	fmt.Println(true && true == true)
	fmt.Println(true && false == false)

	//論理和
	fmt.Println(true || false == true)
	fmt.Println(false || false == true)

	//論理値を反転
	fmt.Println(!true)
	fmt.Println(!false)

}
