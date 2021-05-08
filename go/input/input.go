package main

import "fmt"

// インターフェース
type input interface {
	read() string
}

// 構造体
type keyboard struct{}
type file struct{}
type mic struct{}
type sensor struct{}

func (k keyboard) read() string {
	return "何かキーを押してください"
}

func (f file) read() string {
	return "ファイルパスを指定してください"
}

func (m mic) read() string {
	return "ボタンを押して話してください"
}

func (s sensor) read() string {
	return "計測を開始してください"
}

func main() {
	// データ型を構造体とする input：インターフェース
	methods := []input{
		keyboard{}, file{}, mic{}, sensor{},
	}
	// 実行する
	for i := 0; i < len(methods); i++ {
		fmt.Println(methods[i].read())
	}
}
