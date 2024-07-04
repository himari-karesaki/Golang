package main

import (
	"database/sql"
	"fmt"

	//初期処理で必要なパッケージをインポート
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// データベースに接続
	db, err := sql.Open("mysql", "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました:", err)
		return
	}
	//セットで必ず書く
	defer db.Close()

	// 接続が成功したことを確認
	err = db.Ping()
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました:", err)
		return
	}

	fmt.Println("データベースに接続しました")
}
