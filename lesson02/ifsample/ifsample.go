package main

import "fmt"

func main() {
	score := 100 // 成績

	if score == 100 {
		fmt.Println("満点です！")
	} else if score >= 80 {
		fmt.Println("よくできました！")
	} else if score >= 60 {
		fmt.Println("合格です！")
	} else {
		fmt.Println("赤点です。。。") //複雑なif文を使うことはない　間違っていることを考える必要がある
	}
}
