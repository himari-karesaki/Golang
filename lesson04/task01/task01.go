package main

import "fmt"

func main() {
	//点数を入力する変数
	var score int

	//点数を格納するスライスを作成
	scores := make([]int, 5, 10)

	//無限ループの設定
	for {
		//繰り返し処理
		//「点数を入力してください：」を表示させる
		fmt.Println("点数を入力してください：")
		//点数の入力を受け付ける
		fmt.Scan(&score)
		//「-1」が入力されたら無限ループを抜ける
		if score == -1 {
			break
		}
		//入力した点数をスライスに追加
		scores = append(scores, score)

	}

	//点数の合計値を求める
	for index, value := range scores {
		//繰り返し値が足される処理を行う
	}
	//人数を求める
	count := len(scores)

	// 平均点を求める
	average := 点数の合計値 / count

	float64(average)

	//averageを小数第一位まで四捨五入する関数
	func Round(num float64, pos int) float64 {
		shift := math.Pow10(pos)             // 小数の位置をずらすためのシフト値を算出
		shiftedNum := num * shift            // 四捨五入したい桁を小数第一位にずらす
		roundedNum := math.Round(shiftedNum) // 小数第一位を四捨五入する
		result := roundedNum / shift         // 小数の位置を元に戻す
		return result
	}
	//averageを小数第一位まで四捨五入する
	Round(average, 1)

	//「○○人のテストの平均点は△△点です」と表示させる
	//平均点の出力の際は小数点第一位まで表示させる
	fmt.Printf("%.1f人のテストの平均点は%v点です", count, average)

if length := len(scores);length == 0 {
	//点数が入力されなかった場合の処理
	fmt.Println("点数が入力されませんでした")
}
}
