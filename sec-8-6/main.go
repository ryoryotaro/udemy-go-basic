package main

import "fmt"

//ラベル付きfor

func main() {

	// Loopというラベルの箇所までジャンプしてforを一気に抜ける
	// Loop:
	// 	for {
	// 		for {
	// 			for {
	// 				fmt.Println("START")
	// 				break Loop
	// 			}
	// 			fmt.Println("処理をしない")
	// 		}
	// 		fmt.Println("処理をしない")
	// 	}
	// 	fmt.Println("END")

Loop:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("処理をしない")
	}

}
