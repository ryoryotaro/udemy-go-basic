package main

import "fmt"

//if
//条件分岐

func main() {
	a := 0
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't know")
	}

	//簡易分付きif
	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	x := 0
	if x := 2; true { //ifブロックが優先される
		fmt.Println(x)
	}
	fmt.Println(x)
}
