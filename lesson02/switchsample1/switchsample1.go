package main

import "fmt"

func main() {
	score := 72 // 成績

	//値で判定する場合、範囲での判定はできない10で割ることで10の値で判定する工夫をしている
	switch score / 10 {
	case 10:
		fmt.Println("満点です！")
	case 9, 8:
		fmt.Println("よくできました！")
	case 7, 6:
		fmt.Println("合格です！")
	default:
		fmt.Println("赤点です。。。")
	}
}
