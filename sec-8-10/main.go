package main

import "fmt"

// init 関数
// `main()` よりも先に実行される特殊な関数
func init() {
	fmt.Println("init") // `main()` が実行される前に出力される
}

func main() {
	fmt.Println("Main") // `init()` の後に実行される
}
