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

		//wordsの中にaがはいっているかを確認する処理に変更する
		//wordsの中にaがはいっているかを確認する処理は、for文でwordsの中身を一文字ずつに分解してaが含まれているか確認する
		// for _, word := range words {
		// 	// fmt.Printf("Value: %v\n", word)
		// 	//aが含まれているか確認

		// }

		//endが入力されたらループを終了
		if word == "end" {
			//入力した英単語の表示
			fmt.Println("入力した英単語：")
			for _, word := range words {
				fmt.Println(word)

				//wordがaならcountsのaの値をインクリメントする処理
				//aを変数にする
				//wordの中身を一文字ずつに分解する
				for _, str := range word {
					// fmt.Println(str)
					//ここをどうにかする
					//変数にしたい

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
