package main

import (
	"fmt"
	"strconv"
)

// エラー処理用関数
func plusone(nmstr string) (int, string) {
	res := ": "
	nm, err := strconv.Atoi(nmstr)
	// エラー発生したかを確認する。
	if err != nil {
		res += err.Error()
	} else {
		nm++
		res += "おめでとうございます、" + nmstr
		res += "に1が加算されました"
	}
	return nm, res
}

func main() {
	nm, err := strconv.Atoi("123")
	fmt.Println(nm, err == nil)

	nm, err = strconv.Atoi("ゴー")
	fmt.Println(nm, err == nil)

	nm, res := plusone("53")
	fmt.Println(nm, res)

	nm, res = plusone("What")
	fmt.Println(nm, res)
}
