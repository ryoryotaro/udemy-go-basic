package main

import (
	"fmt"
	"time"
)

// go goroutine
// 並行処理の確認

// 無限ループで "Sub loop" を定期的に出力する関数
func sub() {
	for {
		fmt.Println("Sub loop")
		time.Sleep(100 * time.Millisecond) // 100ms 待機
	}
}

func main() {
	// sub() を goroutine として非同期実行
	// これにより、新しい goroutine が作られ、並行して処理が行われる
	go sub()

	// 無限ループで "Main loop" を定期的に出力
	// `ps -eLf | grep main.go` で確認すると、1つのプロセス内で複数のスレッドが動作していることが分かる
	for {
		fmt.Println("Main loop")
		time.Sleep(200 * time.Millisecond) // 200ms 待機
	}
}
