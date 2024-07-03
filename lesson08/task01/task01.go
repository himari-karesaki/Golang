package main

import (
	"fmt"
	"os"
	"sync"
)

// ファイルの読み込みを行う関数
func readFile(file []string, ch chan<- int, wg *sync.WaitGroup) {

	//file文字数をintにしたいので、int型のスライスを作成
	sum := 0

	// // 文字数のカウントをチャネルに送信
	// ch <- sum
	//文字数を格納する変数
	//ファイルを読み込んで文字数をカウントする
	for i, _ := range file {
		data, err := os.ReadFile(file[i])
		if err != nil {
			fmt.Println("ファイルを読み込めませんでした:", err)
			return
		}
		// 読み込んだデータの文字数をカウント
		// wg.Done()
		sum = len(data)

	}

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
	//readFile関数を実行する
	for i := 0; i < len(files); i++ {
		//goroutineを立ち上げて
		wg.Add(1)
		go readFile(files, ch, &wg)
		fmt.Println(files[i])
	}
	fmt.Println("ゴルーチンの完了を待っています...")

	wg.Wait() // すべてのゴルーチンの完了を待つ

	// チャネルから結果を受信
	sum := <-ch

	fmt.Println("Total sum:", sum)
	fmt.Println("すべてのゴルーチンが終了しました")
}
