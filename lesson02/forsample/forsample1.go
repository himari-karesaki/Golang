package main

import "fmt"

// while文はない
func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			// あまりが0ではない時、つまり奇数の場合、現在の処理をスキップし、次の繰り返しに進む
			continue
		}
		fmt.Println(i)
	}

	//20までの偶数を表示するが、20を超えたら繰り返しから抜ける
	for i := 2; i <= 100; i += 2 {
		if i > 20 {
			// 繰り返しから抜け出す
			break
		}
		fmt.Println(i)
	}
}
