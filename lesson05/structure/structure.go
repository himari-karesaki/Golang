package main

import "fmt"

// 構造体を定義している
// 二つのフィールドを持つPerson型を定義
type Person struct {
	Name string
	Age  int
}

// 構造体にメソッドを定義している
// たい焼きの型Personとたい焼きpとあんこ中身
// 　パーソン型の引数を受け取り、名前と年齢を出力するメソッド
func (p Person) SayHello() {
	fmt.Printf("こんにちは!私の名前は%sです。私は%d歳です。\n", p.Name, p.Age)
}

func (p *Person) IncrementAge() {
	//ポインター渡し
	//中の値でフィールドを使いたいかどうかでポインター渡しを使うか決める
	// main変数の値を変更するためにポインター渡しを使う
	//pの値そのものを操作する
	p.Age++
	fmt.Println("誕生日を迎えました。")
}

func main() {
	p := Person{Name: "山田太郎", Age: 30}
	p.SayHello()     // 出力: こんにちは!私の名前は山田太郎です。私は30歳です。
	p.IncrementAge() // 出力: 誕生日を迎えました。
	p.SayHello()     // こんにちは!私の名前は山田太郎です。私は31歳です。
}
