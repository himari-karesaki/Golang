package main

import (
	"log"
	"os"
)

func main() {
	// ログメッセージを出力
	log.Println("これはログメッセージです")

	// ログの出力先をファイルに変更
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println("ファイルにログを出力します")
}
