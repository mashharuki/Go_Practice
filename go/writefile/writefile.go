package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var input string
	var content string

	for {
		fmt.Print("何か入力してください。's'で保存： ")
		fmt.Scanln(&input)

		if input == "s" {
			break
		}
		content += input
	}

	// バイトデータに変換する。
	wdata := []byte(content)
	err := ioutil.WriteFile("datafile.txt", wdata, 0777)
	// エラーの有無によって分岐する。
	if err != nil {
		fmt.Println("書き込みエラーです。", err)
		os.Exit(1)
	} else {
		fmt.Println(content)
		fmt.Println("以上、ファイルに保存しました。")
	}
}
