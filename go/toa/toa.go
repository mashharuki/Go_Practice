package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 9 * 9
	spfa := fmt.Sprintf("9かける9は%d", 9*9)
	fmt.Println(spfa)
	// 数値と文字列を合体させる(整数を文字列に変換する)
	astr := "9の2乗も" + strconv.Itoa(a)
	fmt.Println(astr)

	b := 3.14159
	spfb := fmt.Sprintf("円周率は大体%.5f", b)
	fmt.Println(spfb)

	bstr := strconv.FormatFloat(b, 'f', 2, 64)
	fmt.Println("もっと大雑把には" + bstr)
}
