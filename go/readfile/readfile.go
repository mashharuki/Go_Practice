package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// ファイルを読み込む
	content, err := ioutil.ReadFile("data.txt")
	// 読み込み結果に応じて分岐する。
	if err != nil {
		fmt.Println("ファイルの読み込みエラー", err)
		os.Exit(1)
	} else {
		cstr := string(content)
		lows := strings.Split(cstr, "¥n")
		// ループする。
		for i, v := range lows {
			if i > 0 && len(v) > 1 {
				data := strings.Split(v, ",")
				fmt.Printf("%sは%s円¥n", data[0], data[1])
			}
		}
	}
}
