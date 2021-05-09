package main

import "fmt"

func counter() func() {
	ctr := 0
	fmt.Println("カウンタを初期化しました")
	fmt.Printf("カウンタのアドレスは %p¥n", &ctr)
	// 戻り値
	return func() {
		ctr++
		fmt.Printf("カウンタの値は %d¥n", ctr)
		fmt.Printf("カウンタのアドレスは %p¥n", &ctr)
	}
}

func main() {
	countfnc := counter()
	countfnc()
	countfnc()
	countfnc()
}
