package main

import (
	"fmt"
	"os"
)

// defer の動作確認

// TestDefer は defer の基本的な動作を示す関数
// defer された fmt.Println("END") は関数終了時に実行される
func TestDefer() {
	defer fmt.Println("END") // 関数の最後に実行される
	fmt.Println("START")     // これは即時実行
}

// RunDefer は複数の defer の実行順序を確認する関数
// defer は後入れ先出し (LIFO) で実行される
func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3") // 最後に登録されたため最初に実行される
}

func main() {
	// defer の基本動作を確認
	TestDefer()

	// 無名関数を defer で登録
	// defer は関数終了時に実行されるため、A → B → C の順で出力される
	defer func() {
		fmt.Println("A")
		fmt.Println("B")
		fmt.Println("C")
	}()

	// defer の実行順序を確認
	RunDefer()

	// ファイルを作成し、defer で確実に Close() する
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() // 関数終了時に必ずファイルを閉じる

	// ファイルにデータを書き込む
	file.Write([]byte("Hello"))
}
