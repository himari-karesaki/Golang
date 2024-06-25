package main

import "fmt"

func main() {
	//{}は配列に入れる値の数を指定している
	numbers := [5]int{50, 12, 39, 77, 65}
	//indexと値がこの順番で取得できるようになっている _を使うとindexを無視できる
	for index, value := range numbers {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}
}
