package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	//GORMの読み込み
	"gorm.io/gorm"
)

type Book struct {
	ID        uint `gorm:"primaryKey"` //GORMの主キーであることを示している
	Title     string
	Price     int
	CreatedAt time.Time
}

func main() {
	// データベースに接続
	dsn := "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //&gorm.Config{}はGORMの設定情報があれば入れる
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました:", err)
	}
	//GORMは自分でクローズする
	//ユーザーがクローズする必要がない

	// テーブルを自動的に作成
	err = db.AutoMigrate(&Book{}) //Bookの構造体を作って、この中に入れれるようにする
	if err != nil {
		fmt.Println("データベースのマイグレーションに失敗しました:", err)
	}

	// レコードを挿入
	// 複数名から単数名に変更する
	books := []Book{
		{Title: "はじめてのMySQL", Price: 2980},
		{Title: "はじめてのGo", Price: 1980},
		{Title: "はじめてのHTML", Price: 1000},
		{Title: "はじめてのCSS", Price: 1000},
	}

	//bookスライスのレコードを全て入れている
	result := db.Create(&books)
	if result.Error != nil {
		fmt.Println("レコードの挿入に失敗しました")
	}
	fmt.Printf("%d レコード挿入しました。\n", result.RowsAffected) //影響があった件数が返ってくる　決まっている戻り値
}
