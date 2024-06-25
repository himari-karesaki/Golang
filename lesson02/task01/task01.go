package main

import "fmt"

func main() {
	var name string = "山田太郎"
	var age int = 25
	var from string = "東京"
	var fan string = "読書"

	fmt.Printf("はじめまして、%vと申します。\n年齢は%v歳で、%v都出身です\n趣味は%vです。\nよろしくお願いします。", name, age, from, fan)
}
