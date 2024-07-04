package main

import (
	"fmt"
	"os"
	"sync"
)

// ファイルの読み込みを行う関数
func readFile(file string, ch chan<- int, wg *sync.WaitGroup) {
	//引数で指定されたファイルを読み込み
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("ファイルを読み込めませんでした:", err)
		return
	}
	// 読み込んだデータの文字数をカウント
	sum := len(data)
	//文字数のカウントをチャネルに送信
	ch <- sum
	wg.Done()
}

func main() {
	//チャネルに送信された値のスライスの作成
	ch := make(chan int)
	//"file1.txt", "file2.txt", "file3.txt" が要素のスライスを作成
	files := []string{"./file1.txt", "./file2.txt", "./file3.txt"}
	// WaitGroupを作成
	var wg sync.WaitGroup

	totalSum := 0
	//readFile関数を実行する
	for i := 0; i < len(files); i++ {
		wg.Add(1)
		file := files[i]
		go readFile(file, ch, &wg) // ファイルの読み込みをゴルーチンで実行
		totalSum += <-ch
	}

	wg.Wait() // すべてのゴルーチンの完了を待つ

	// 文字数の合計を表示
	fmt.Printf("合計文字数：%v", totalSum) //
}
