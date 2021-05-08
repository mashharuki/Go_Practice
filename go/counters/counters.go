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

func main() {
	// counter関数を一度だけで呼び出す。
	countfnc := counter()
	// 戻り値を表示する。
	countfnc()
	countfnc()
	countfnc()
}
