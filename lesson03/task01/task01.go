package main

import "fmt"

func main() {
	var subtotal float64 = 1500                  //商品の原価
	taxPrice := tax(subtotal)                    //消費税率
	var taxProduct float64 = subtotal + taxPrice //商品の税込価格

	fmt.Printf("%v円の商品の税込価格は%v円(消費税は%v円)です\n", subtotal, taxProduct, taxPrice)
}

func tax(price float64) float64 {
	var taxRate float64 = 0.1 //消費税率
	return price * taxRate
}
