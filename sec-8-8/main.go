package main

import "fmt"

// panic と recover の動作確認

func main() {
	// defer を使い、関数終了時に panic をキャッチする無名関数を実行
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("Recovered from panic:", x) // panic の内容を出力
		}
	}()

	panic("runtime error") // 明示的に panic を発生させる

	fmt.Println("Start") // panic により実行されない
}
