package main

import (
	"fmt"
	"math"
)

func main() {
	//点数を入力する変数(intは符号付き整数)
	var score int

	//点数を格納するスライスを作成
	//容量の最大値はあえて指定しない
	scores := make([]int, 0)

	//無限ループの設定
	for {
		//「点数を入力してください：」を表示させる
		fmt.Print("点数を入力してください：")
		//点数の入力を受け付ける
		fmt.Scan(&score)

		//「-1」が入力されたら無限ループを抜ける
		if score == -1 {

			if length := len(scores); length == 0 {
				//点数が入力されなかった場合の処理
				fmt.Println("点数が入力されませんでした")
			}
			break
		}

		//入力した点数をスライスに追加
		scores = append(scores, score)

		// var newAverage float64 = float64(average)

		// 「○○人のテストの平均点は△△点です」と表示させる
		// 平均点の出力の際は小数点第一位まで表示させる
		// fmt.Printf("%v人のテストの平均点は%.1f点です\n", count, newAverage)
	}
	// 点数の合計値を求める
	var sum float64 = sumScore(scores)
	// 人数を求める処理
	count := len(scores)
	// 平均点を求める処理
	average := sum / float64(count)

	//averageを小数第一位まで四捨五入する
	Round(average, 1)

	newAverage := float64(sumScore(scores) / float64(len(scores)))

	// 「○○人のテストの平均点は△△点です」と表示させる
	// 平均点の出力の際は小数点第一位まで表示させる
	fmt.Printf("%v人のテストの平均点は%.1f点です\n", count, newAverage)
}

// 点数の合計値を求める関数
func sumScore(numbers []int) float64 {
	sum := 0
	// var number int
	for _, number := range numbers {
		sum += number
	}
	return float64(sum)
}

// averageを小数第一位まで四捨五入する関数
func Round(num float64, pos int) float64 {
	shift := math.Pow10(pos)             // 小数の位置をずらすためのシフト値を算出
	shiftedNum := num * shift            // 四捨五入したい桁を小数第一位にずらす
	roundedNum := math.Round(shiftedNum) // 小数第一位を四捨五入する
	result := roundedNum / shift         // 小数の位置を元に戻す
	return result
}
