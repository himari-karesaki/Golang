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
	counts := make(map[rune]int)
	// counts['a'] = 1
	for {
		//aの個数をカウントする
		counts['a']++
		fmt.Println(counts['a'])

		if value, ok := counts['a']; ok {
			fmt.Printf("a: %d\n", value)
			break // key1 exists. value = 10
		}

	}

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
				//この中でワードを一文字ずつに分解する処理を書く

			}

			//countsにaが含まれていたらカウントする
			counts['a']++

			// counts["a"]++
			//mapcountsのaの値を表示する
			fmt.Println(counts['a'])

			fmt.Println("アルファベットごとの文字数：")
			//アルファベットごとに文字数を表示する処理
			for _, value := range counts {
				fmt.Printf("a:%d\n", value)
				//ここに問題がありそう

				// fmt.Println(count)
			}

			break
		}

		//入力した値をスライスwordsに格納
		words = append(words, word)
		// fmt.Println(words)

	}
}
