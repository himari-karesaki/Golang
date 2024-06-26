package main

import "fmt"

func main() {
	number1 := 100
	number2 := 10
	sumMethod1(number1, number2) //number1とnum1のように変わっていても良い 関数の定義はどこに書いてもよい
	// ----- 追記（ここから）-----
	result2 := sumMethod2(number1, number2)
	fmt.Println("sumMethod2の結果は", result2)
	// ----- 追記（ここまで）-----
}

// 関数群
func sumMethod1(num1 int, num2 int) {
	result := num1 + num2
	fmt.Println("sumMethod1の結果は", result)
}

// ----- 追記（以下の関数）-----
func sumMethod2(num1 int, num2 int) int { //戻り値がある場合はintのように指定する
	result := num1 + num2
	return result
}
