package main

import "fmt"

func main() {
	// 無名関数を変数に代入funcの後に名前がない関数のこと
	add := func(a, b int) int {
		return a + b
	}

	result := add(3, 5)
	fmt.Println(result) // 8
}
