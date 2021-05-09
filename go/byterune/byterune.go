package main

import "fmt"

func main() {
	str := "internationalzation"
	strja := "だるまさんが転んだ"
	// 数値データに変換
	bt := []byte(str)
	fmt.Println(bt)
	// 数値データに変換
	rn := []byte(strja)
	fmt.Println(rn)

	// 文字列データに戻す
	strback := ""
	for i := 0; i < len(bt); i++ {
		strback += string(bt[i])
	}

	strbackja := ""
	for i := 0; i < len(rn); i++ {
		strbackja += string(rn[i])
	}

	fmt.Println(strback)
	fmt.Println(strbackja)
}
