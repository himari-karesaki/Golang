package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	ID        int
	Title     string
	Price     int
	CreatedAt time.Time
}

func main() {
	var title string

	fmt.Print("検索するタイトルを入力してください: ")
	fmt.Scan(&title)

	// データベースに接続
	db, err := sql.Open("mysql", "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました:", err)
		return
	}
	defer db.Close()

	// クエリを実行
	var dbTitle string = title

	var book Book

	err = db.QueryRow("SELECT * FROM books WHERE title = ?", dbTitle).Scan(&book.ID, &book.Title, &book.Price, &book.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("指定されたタイトルのレコードが見つかりませんでした")
		} else {
			fmt.Println("クエリの実行に失敗しました:", err)
		}
		return
	}

	// 結果を出力
	fmt.Println(book.ID, book.Title, book.Price, book.CreatedAt)
}
