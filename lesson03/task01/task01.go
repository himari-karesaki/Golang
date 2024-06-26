package main

import "fmt"

func main() {
	var subtotal float64 = 1500                  //商品の原価
	taxPrice := tax()                            //消費税額
	var taxProduct float64 = subtotal + taxPrice //商品の税込価格

	fmt.Printf("%v円の商品の税込価格は%v円(消費税は%v円)です。\n", subtotal, taxProduct, taxPrice)
}

// 消費税を計算する関数
func tax() float64 {
	var price float64 = 1500  //商品の原価
	var taxRate float64 = 0.1 //消費税率
	return price * taxRate
}
