package main

import (
	"fmt"
	"os"

	"unicode/utf8"
)

// file1:="./file1.txt"
// file2:="./file2.txt"
// file3:="./file3.txt"

func readFile() {
	//引数で指定したファイルを読み込む処理
	//file1の読み込み
	// ReadFile関数
	file1, err := os.ReadFile("file1.txt")
	if err != nil {
		fmt.Println("ファイルを読み込めませんでした:", err)
		return
	}
	fmt.Println(utf8.RuneCountInString(string(file1)))

	//file2の読み込み
	file2, err := os.ReadFile("file2.txt")
	if err != nil {
		fmt.Println("ファイルを読み込めませんでした:", err)
		return
	}
	fmt.Println(utf8.RuneCountInString(string(file2)))

	//file3の読み込み
	file3, err := os.ReadFile("file3.txt")
	if err != nil {
		fmt.Println("ファイルを読み込めませんでした:", err)
		return
	}
	fmt.Println(utf8.RuneCountInString(string(file3)))
	//内容の文字数を計算する処理
	// count := 0
	//文字数のカウントのスライス
	//content1{a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z}
	// for _, word1 := range content1 {
	// 	int(word1)
	// 	count(word1)++
	// 	return word1
	// }
	// for _, word2 := range content2 {
	// 	count++
	// 	return string(word2)
	// }
	//結果をチャネル経由でmain関数に送信する処理
}

func main() {
	//スライスの作成
	//"file1.txt", "file2.txt", "file3.txt" が要素

	//goroutineを立ち上げてreadFile関数を実行する
	readFile()
	//結果を受信するチャネルを作成
	// ch := make(chan int)

	//
	// go readFile(content1, ch)
	// go readFile(context2, ch)
	// readFile()
	//file1
	//file2
	//file3

	//main関数からチャネル��由で受信する
	//結果を表示する

	//waitGroupを使って全てのgoroutineが終了するのを待つ

	//file1のGoroutineから文字数を受け取る
	//file2のGoroutineから文字数を受け取る
	//file3のGoroutineから文字数を受け取る

	// for _, v := range files {
	// 	//合計を計算する処理
	// 	count++

	// 	//最終的な文字数を表示する
	// 	fmt.Println(count)
	// }

}
