package main

import "fmt"

func processData(data string) {
	if len(data) == 0 {
		panic("データが空です") //panicはプログラムを続けることができない致命的なエラーを示す
	}
	fmt.Println("データの長さ:", len(data))
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("パニックが発生しました:", r) //recoverはpanicが発生したときに呼び出される
		}
	}() //関数が実行される直前で実行される

	processData("こんにちは")
	processData("")
	fmt.Println("プログラムが終了しました")
}
