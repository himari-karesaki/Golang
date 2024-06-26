package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Print("割られる数を入力してください: ")
	_, err := fmt.Scan(&a)

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

	//割る数についての処理
	fmt.Print("割る数を入力してください: ")

	_, err = fmt.Scan(&b)

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
