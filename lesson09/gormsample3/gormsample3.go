package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	ID        int
	Title     string
	Price     int
	CreatedAt time.Time
}

func main() {
	// データベースに接続
	dsn := "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました:", err)
		return
	}

	// クエリを実行
	var id int = 2
	var book Book
	result := db.First(&book, id) //Firstで一件だけ取得する2件以上ある場合は最初の1件が取得される1件が保証されているのはFirst　二件以上はFindを使って Whereを使って条件を書くこともできる
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("指定されたIDのレコードが見つかりませんでした")
		} else {
			fmt.Println("クエリの実行に失敗しました:", result.Error)
		}
		return
	}

	// 結果を出力
	fmt.Println(book.ID, book.Title, book.Price, book.CreatedAt)
}
