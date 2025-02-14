package main

import (
	"fmt"
	"strconv"
)

func main() {

	//型
	//整数型

	var i int = 400
	// 64bitのPCだと、暗黙的にint64になる（環境依存）

	//明示的にしてすることも可能
	var i2 int64 = 200

	fmt.Println(i + 50)

	//これは型が合わなくてエラーになる
	//明示的な方と暗黙的に変換された型は駄目らしい
	//fmt.Println(i + i2)

	//変数の型を表示させる
	fmt.Printf("%T\n", i2)

	//型を変換
	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i + int(i2))

	//型
	//浮動小数点型
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	//これはfloat64になる（環境依存ではない）
	fl := 3.2

	fmt.Println(fl + fl64)
	fmt.Printf("%T,%T\n", fl64, fl)

	var fl32 float32 = 1.3
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	//Inf（無限大）: 浮動小数点の範囲外の大きな数を表す。
	//例えば、0で割るような計算（1.0 / 0.0）は正の無限大（+Inf）を返し、
	//負の数を0で割ると負の無限大（-Inf）を返す。
	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	//NaN（Not a Number）: 計算の結果が無効な場合に返される特別な値。
	//例えば、0.0を0.0で割るなど、定義できない計算はNaNになる。

	/*
		ちなみに、intではゼロ除算はエラーになるが、浮動小数点ではゼロ除算を許容している
		理由はゼロ除算が発生したとしても、プログラム全体がクラッシュせずに計算を継続できるようにするため。
		たとえば、シミュレーションや数値解析のような計算では、途中でゼロ除算のような未定義な操作が起きる可能性があるらしいが、
		そこでプログラムを停止してしまうと全体の計算が台無しになってしまう。
		IEEE 754規格に従うことで、ゼロ除算が発生した場合も、適切に無限大（+Infや-Inf）やNaN（Not a Number）という特殊な値を返し、
		例外処理を行わずに計算を続行することができる。これにより、計算結果がすぐに分からなくなることを防ぎ、エラー処理の柔軟性が高まる。
	*/
	nan := zero / zero
	fmt.Println(nan)

	//型
	//論理値型
	var t, f bool = true, false
	fmt.Println(t, f)

	//型
	//文字列型
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	//複数行表示
	fmt.Println(`test
	test
		test`)

	//エスケープする
	fmt.Println("\"")
	fmt.Println(`"`)

	//文字列はByte型の集まり
	fmt.Println(string(s[0]))

	//型
	//byte(uint8)型

	byteA := []byte{72, 73}
	fmt.Println(byteA)

	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)

	fmt.Println(string(c))

	//配列型
	//あとから要素数を変更できない。可変長の配列はスライスという
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"} //要素数を省略
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	//値の取り出し
	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	fmt.Println(arr2[2-1])

	//値の操作
	arr2[2] = "C"
	fmt.Println(arr2)

	//var arr5 [4]int
	//arr5 = arr1 //これはエラーになる。要素数の違うもの同士は代入できない
	//fmt.Println(arr5)

	//要素数を調べる関数
	fmt.Println(len(arr1))

	//interface & nil
	//すべての型と互換性がある
	//下の例だと、xにはすべての値を代入することができる
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	x = 3.14
	fmt.Println(x)

	x = "A"
	fmt.Println(x)

	x = [3]int{1, 2, 3}
	fmt.Println(x)

	//ただしこれは以下のようなエラーになる
	// invalid operation: x + 1 (mismatched types interface{} and int)
	//型x = 2
	//fmt.Println(x + 1)

	//型変換
	var i_2 int = 1
	fl64_2 := float64(i_2)
	fmt.Println(fl64_2)

	i_3 := int(fl64_2)
	fmt.Printf("i = %T\n", i_3)

	fl_3 := 10.5
	i3 := int(fl_3)
	fmt.Println(i3) //10になる（小数点以下は切り捨て）

	var s_22 string = "100"
	// これは2つ目の戻り地を _ に入れて破棄するやり方
	i_22, _ := strconv.Atoi(s_22)
	fmt.Printf("i_22 = %T\n", i_22)

}
