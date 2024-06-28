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

	fmt.Println("英単語を入力してください(endと入力するとループを終了します): ")

	//無限ループの作成
	for {
		//英単語の標準入力
		fmt.Scan(&word)

		//endが入力されたらループを終了
		if word == "end" {
			//入力した英単語の表示
			fmt.Println("入力した英単語：")
			for _, word := range words {
				fmt.Println(word)

				//wordの中身を一文字ずつに分解する
				for _, str := range word {
					counts[str]++
				}

			}

			fmt.Println("アルファベットごとの文字数：")
			// fmt.Println(counts)
			for key, count := range counts {
				fmt.Printf("%v:%d\n", string(key), count)
			}
			break
		}

		//入力した値をスライスwordsに格納
		words = append(words, word)
	}
}
