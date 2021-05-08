package main

import "fmt"

// @return func()
func counter() func() {
	ctr := 0
	fmt.Println("カウンタを初期化する。")
	return func() {
		ctr++
		fmt.Println(ctr)
	}
}

// @return func(int)
func freecounter(start int) func(int) int {
	ctr := start
	fmt.Printf("フリーカウンターを%dから始めますか？¥n", ctr)
	// 戻り値
	return func(add int) int {
		fmt.Printf("%dを足して", add)
		ctr += add
		return ctr
	}
}

func main() {
	// freecounter関数を一度だけで呼び出す。
	freecnt := freecounter(10)
	// 戻り値を表示する。
	fmt.Println(freecnt(2))
	fmt.Println(freecnt(5))
	fmt.Println(freecnt(7))
}
