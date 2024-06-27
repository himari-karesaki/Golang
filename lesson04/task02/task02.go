package main

import (
	"fmt"
)

func main() {
	//英単語を入力する変数の宣言
	var word string
	//英単語を格納するスライスの宣言
	words := make([]string, 0)
	//アルファベットの文字数を格納する変数の宣言
	// counts := make(map[rune]int)
	// counts['a']++
	// fmt.Println(counts)

	fmt.Print("英単語を入力してください(endと入力するとループを終了します): ")

	//無限ループの作成
	for {
		//英単語の標準入力
		fmt.Scan(&word)

		// //mapに文字数を格納する
		// counts := make(map[rune]int)
		// fmt.Println(counts)

		//endが入力されたらループを終了
		if word == "end" {
			//入力した英単語の表示
			fmt.Println("入力した英単語：")
			for _, word := range words {
				fmt.Println(word)
			}
			break
		}

		//入力した値をスライスwordsに格納
		words = append(words, word)
		fmt.Println(words)

	}

	// マップの要素aをカウント
	//もし文字aが入っていたらカウントする
	// var a string
	// if value, exists := counts["a"]; exists {
	// 	counts["a"]++
	// 	fmt.Println("a:", value)
	// } else {
	// 	fmt.Println("a does not exist")
	// }

}

// //アルファベットの文字数をカウントする関数
// func countA(a []string) int {
// 	for _, a := range word {
// 	counts["a"]++
// }

// func countWords(words []string) int {
// 	// キーの存在チェック
// 	if value, exists := m["a"]; exists {
// 		fmt.Println("a:", value)
// 	}
// }
