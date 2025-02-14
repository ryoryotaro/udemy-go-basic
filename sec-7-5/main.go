package main

import "fmt"

//関数
//クロージャーの実装（ある関数が外の変数を覚えて、その変数を使い続けることができる仕組み）

/*
クロージャー（closure）とは、関数とその関数が定義された環境（変数など）を一緒に閉じ込めたものを指します。つまり、関数が外部の変数を覚えていて、
その関数の中からその変数にアクセスし続けることができる、という仕組みです。

クロージャーの特徴

    関数の外で定義された変数にアクセスできる。
    その変数の値を保持し、変更を追跡できる。
    外部関数の実行が終了しても、その変数が消えることはなく、後から呼び出してもその状態を保持したままである。
*/

func Later() func(string) string {
	var store string // 外部変数 `store` を定義
	return func(next string) string {
		s := store   // 外部変数 `store` にアクセス
		store = next // `store` を更新
		return s     // 前の値を返す
	}

}

func main() {
	/*
		最初の Later() の呼び出しで返されるのは、無名関数そのものです。
		Later の中で定義された store 変数にアクセスし続ける「状態を持った関数」が f に代入されます。
	*/
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("My")) //2回目以降の呼び出しは return された無名関数（クロージャー）だけが実行されます。
	fmt.Println(f("name"))
	fmt.Println(f("is"))
	fmt.Println(f("golang !"))

}

/*
それぞれの fmt.Println(f(...)) の動作:

    最初の f("Hello") では、まだ store は空なので、空文字列 ("") を返します。次に store に "Hello" が保存されます。
    次に f("My") では、store に "Hello" が入っているので "Hello" を返し、store は "My" に更新されます。
    同様に、f("name") では "My" を返し、store は "name" に更新されます。
    f("is") では "name" を返し、store は "is" に更新されます。
    最後の f("golang !") では "is" を返し、store は "golang !" に更新されます。

結果として、store に一時的に保持されていた前回の値が出力されるという動作になります。
*/
