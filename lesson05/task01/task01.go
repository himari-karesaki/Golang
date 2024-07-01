package main

import "fmt"

// Petインターフェースの定義
type Pet interface {
	Name() string //ペットの名前を文字列で返す
	Eat()         //ペットが食事をすることを表す
	Play()        //ペットが遊ぶことを表す
	Sleep()       //ペットが眠ることを表す
}

// 構造体の定義
// Dog構造体
type Dog struct {
	name   string
	energy int
}

// Dog関連のメソッド
// Dogの名前を返すメソッド
func (d *Dog) Name() string {
	return d.name
}

// DogのEatメソッドでエネルギーが10増えた値を返す
func (d *Dog) Eat() int {
	//Dogのenergyを10増やす処理
	return d.energy + 10
}

// DogのPlayメソッドでエネルギーが5減った値を返す
func (d *Dog) Play() int {
	//Dogのenergyを5減らす処理
	return d.energy - 5
}

// DogのSleepメソッドでエネルギーが20増えた値を返す
func (d *Dog) Sleep() int {
	//Dogのenergyを20増やす処理
	return d.energy + 20
}

// Cat構造体
type Cat struct {
	name string
	mood string
}

// Cat関連のメソッド
// Catの名前を返すメソッド
func (c *Cat) Name() string {
	return c.name
}

// CatのEatメソッドで文字列"happy"を返す
func (c *Cat) Eat() string {
	//Catのmoodをhappyにする処理
	return c.mood
}

// CatのPlayメソッドで文字列"playful"を返す
func (c *Cat) Play() string {
	//Catのmoodをplayfulにする処理
	return c.mood
}

// CatのSleepメソッドで文字列"sleepy"を返す
func (c *Cat) Sleep() string {
	//Catのmoodをsleepyにする処理
	return c.mood
}

func main() {
	//Petスライスを作成
	Pet := []any{
		Dog{name: "バディ", energy: 50},
		Cat{name: "ウィスカー", mood: ""},
	}

	TotalAreaAndPerimeter(Pet)
}

func TotalAreaAndPerimeter(pets []any) {
	// var name, mood string
	// var energy int

	for _, pet := range pets {
		switch v := pet.(type) {
		case Dog:
			name := v.Name()
			eat := v.Eat()
			play := v.Play()
			sleep := v.Sleep()
			fmt.Printf("%vの情報：\n", name)
			fmt.Printf("%vは食事をしています。エネルギー：%v\n", name, eat)
			fmt.Printf("%vは遊んでいます。エネルギー：%v\n", name, play)
			fmt.Printf("%vは眠っています。エネルギー：%v\n", name, sleep)

		case Cat:
			name := v.Name()
			mood := v.mood()
			fmt.Printf("%vの情報：\n", name)
			fmt.Printf("%vは食事をしています。気分： %v\n", name, mood)
			fmt.Printf("%vは遊んでいます。気分： %v\n", name, mood)
			fmt.Printf("%vは眠っています。気分： %v\n", name, mood)
		}
	}
	return
}
