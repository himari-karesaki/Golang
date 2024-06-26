package main

import "fmt"

func main() {
	// Print関数
	fmt.Print("Hello, World!\n")

	// Print関数：複数の値を渡す例
	fmt.Println("Hello,", "World!")

	// Printf関数
	name := "Alice"
	age := 25
	fmt.Printf("%sは%d歳です。\n", name, age)

	// Scan関数
	fmt.Print("名前を入力してください: ")
	fmt.Scan(&name)
	fmt.Print("年齢を入力してください: ")
	fmt.Scan(&age)
	fmt.Printf("%sは%d歳です。\n", name, age)
}
