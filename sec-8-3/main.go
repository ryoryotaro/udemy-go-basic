package main

import "fmt"

//for
//繰り返し

func main() {
	/*
		i := 0
		for {
			i++
			if i == 3 {
				break
			}
			fmt.Println("loop")
		}
	*/
	/*
		point := 0
		for point < 10 {
			fmt.Println(point)
			point++
		}
	*/
	/*
		for i := 0; i < 10; i++ {
			if i == 3 {
				continue //forをSKIPする（3が表示されない）
			}
			if i == 6 {
				break

			}
			fmt.Println(i)
		}
	*/
	/*
		arr := [3]int{1, 2, 3}
		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i])
		}
	*/
	/*
		arr := [3]int{1, 2, 3}

		for i, v := range arr {
			fmt.Println(i, v) //indexの番号と値を出す
		}
	*/

	/*
		sl := []string{"Python", "PHP", "GO"} //要素数が可変なのをスライスと呼ぶ
		for i, v := range sl {
			fmt.Println(i, v)
		}
	*/

	m := map[string]int{"apple": 100, "banana": 200} //key:valueで値が入る配列をMAPという
	for k, v := range m {
		fmt.Println(k, v)
	}

}
