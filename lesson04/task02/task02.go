package main

import "fmt"

func main() {
	//英単語を入力する変数の宣言
	var word string
	//英単語を格納するスライスの宣言
	words := make([]string, 0)
	//アルファベットの文字数を格納する変数の宣言
	counts := make(map[rune]int)

	fmt.Print("英単語を入力してください(endと入力するとループを終了します): ")

	//無限ループの作成
	for {
		//英単語の標準入力
		fmt.Scan(&word)
		//endが入力されたらループを終了
		if word == "end" {
			break
		}

		//入力した値をスライスwordsに格納
		words = append(words, word)
		fmt.Println(words)

		
	}
}

//アルファベットの文字数をカウントする関数
func countA(a []string) int {
	for _, a := range word {
	counts["a"]++
}

// func countWords(words []string) int {
// 	// キーの存在チェック
// 	if value, exists := m["a"]; exists {
// 		fmt.Println("a:", value)
// 	}
// }