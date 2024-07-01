package main

import (
	"fmt"
	"math"
)

// Rectangle 構造体は、長方形を表します
type Rectangle struct {
	Width  float64 // 長方形の幅
	Height float64 // 長方形の高さ
}

// Area メソッドは、長方形の面積を計算します
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter メソッドは、長方形の周囲の長さを計算します
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 構造体は、円を表します
type Circle struct {
	Radius float64 // 円の半径
}

// Area メソッドは、円の面積を計算します
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter メソッドは、円の周囲の長さ（円周）を計算します
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// TotalAreaAndPerimeter 関数は、any型のスライスを受け取り、面積と周囲の長さの合計を計算します
func TotalAreaAndPerimeter(shapes []any) (float64, float64) {
	var totalArea, totalPerimeter float64
	for _, shape := range shapes {
		//switch shape.(type)で型アサーションを行い、型によって処理を分岐します
		//goの外の世界にある型を知るためには、型アサーションを使うjson型を持ってくる時などに使う　値を受け取るときに型がわからない時に使う
		switch v := shape.(type) {
		case Rectangle:
			area := v.Area()
			perimeter := v.Perimeter()
			fmt.Printf("図形: Rectangle, 面積: %.2f, 周囲の長さ: %.2f\n", area, perimeter)
			totalArea += area
			totalPerimeter += perimeter
		case Circle:
			area := v.Area()
			perimeter := v.Perimeter()
			fmt.Printf("図形: Circle, 面積: %.2f, 周囲の長さ: %.2f\n", area, perimeter)
			totalArea += area
			totalPerimeter += perimeter
		}
	}
	return totalArea, totalPerimeter
}

func main() {
	// 図形のスライスを作成し、長方形と円を追加します
	shapes := []any{
		Rectangle{Width: 5, Height: 10},
		Circle{Radius: 7},
	}

	// 図形の面積と周囲の長さの合計を計算します
	totalArea, totalPerimeter := TotalAreaAndPerimeter(shapes)

	// 結果を出力します
	fmt.Printf("面積の合計: %.2f\n", totalArea)
	fmt.Printf("周囲の長さの合計: %.2f\n", totalPerimeter)
}
