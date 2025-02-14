package main

import "fmt"

// 型スイッチ (type switch) を使った型判定関数
func anything(a interface{}) {
	// `a` の型を判定して適切な処理を行う
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?") // 文字列なら "!?" を追加して出力
	case int:
		fmt.Println(v + 1000) // 数値なら 1000 を加算して出力
	}
}

func main() {
	// `anything` 関数を呼び出し、型ごとの動作を確認
	anything("aaa") // "aaa!?" を出力
	anything(1)     // 1001 を出力

	// `interface{}` 型の変数を用意
	var x interface{} = "3"

	// x.(int) は型が合わないため、実行時エラー（panic）が発生するためコメントアウト
	// i := x.(int)
	// fmt.Println(i + 2)

	// 型アサーション（成功可否をチェック）
	f, isFloat64 := x.(float64) // x を float64 型として取り出そうとする
	fmt.Println(f, isFloat64)   // 変換失敗 → 0 false

	// 型アサーションの安全なチェック
	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, "x is string") // "3 x is string" を出力
	} else {
		fmt.Println("I don't know.")
	}

	// 型スイッチを使った型判定
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string") // "string" を出力
	default:
		fmt.Println("I don't know.")
	}

	// 型スイッチ（変数を受け取るバージョン）
	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string") // "3 string" を出力
	}
}
