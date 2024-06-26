package main

import (
	"fmt"
	// "strconv"
)

func main() {
	var a, b int

	fmt.Print("割られる数を入力してください: ")
	_, err := fmt.Scan(&a)

	// // a, err = strconv.Atoi(a)
	// //割られる数についての処理

	// // a, err = strconv.Atoi(str)
	// // a, err = strconv.Atoi(a)
	// //入力値の例
	// // str1 := "12"

	// a, err = strconv.Atoi(str1) //str1にスキャンしたaを代入したい
	if err != nil {
		//入力値が数値でなかった場合のエラー処理
		fmt.Println("エラー:数値を入力してください")
		fmt.Println("処理を終了します")
		return
	} else if a == 0 {
		//aに0が入力された場合の処理
		println("エラー:0で割り算しないでください")
		println("処理を終了します")
		return
	}

	// _, err := fmt.Scan(&b)
	//割る数についての処理
	fmt.Print("割る数を入力してください: ")
	// fmt.Scan(&b)
	_, err = fmt.Scan(&b)
	// println(b)
	// b, err = strconv.Atoi(str2)
	if err != nil {
		//エラー処理
		fmt.Println("数値を入力してください")
		fmt.Println("処理を終了します")
		return
	} else if b == 0 {
		//bに0が入力された場合の処理
		println("エラー:0で割り算しないでください")
		println("処理を終了します")
		return
	}
	// println("こんばんは")
	c := a / b
	// println("さようなら")
	fmt.Printf("%d ÷ %d = %d\n", a, b, c)
	fmt.Println("処理を終了します")
}
