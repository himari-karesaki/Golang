package main

import "fmt"

// Petインターフェースの定義
type Pet interface {
	Name()  //ペットの名前を文字列で返す
	Eat()   //ペットが食事をすることを表す
	Play()  //ペットが遊ぶことを表す
	Sleep() //ペットが眠ることを表す
}

// 構造体の定義
// Dog構造体
type Dog struct {
	name   string
	energy int
}

// Cat構造体
type Cat struct {
	name string
	mood string
}

func (d Dog) Name() string {
	return d.name
}

func (d Dog) Eat() int {
	//Dogのenergyを10増やす処理
	return d.energy + 10
}

func (d Dog) Play() int {
	//Dogのenergyを5減らす処理
	return d.energy - 5
}

func (d Dog) Sleep() int {
	//Dogのenergyを20増やす処理
	return d.energy + 20
}

func (c Cat) Name() string {
	return c.name
}

func (c Cat) Eat() string {
	//Catのmoodをhappyにする処理
	return "happy"
}

func (c Cat) Play() string {
	//Catのmoodをplayfulにする処理
	return "playful"
}

func (c Cat) Sleep() string {
	//Catのmoodをsleepyにする処理
	return "sleepy"
}

// Dog 関数は、any型のスライスを受け取り、DogなのかCatなのか判断し、どちらかの挙動を返す
func DogOrCat(Pets []any) (string, int) {
	for _, pet := range Pets {
		switch v := pet.(type) {
		case Dog:
			name := v.Name()
			energyCheck := v.Eat()
			fmt.Printf("バディの情報:\nバディは%v。 エネルギー： %v\n", name, energyCheck)

		case Cat:
			name := v.Name()
			moodCheck := v.Sleep()
			fmt.Printf("ウィスカーの情報:\nウィスカーは%v。 気分： %v\n", name, moodCheck)
		}
	}
	return "", 0
}

// main関数
func main() {
	//Petスライスを作成
	Pets := []any{
		Dog{name: "ポチ", energy: 50},
		Cat{name: "ネコ", mood: "sleepy"},
	}

	//Petスライスの各要素に対し、各メソッドを実行する処理

	//結果を出力する処理
	//あとで動作のところを追加
	fmt.Printf("%v%v", DogOrCat(Pets))
}
