package main

import "fmt"

func main() {
	normalStr := "Go言語でLet's go!"
	emptyStr := ""
	emptyStr += "昔は空でした"
	fmt.Println(normalStr)
	fmt.Println(emptyStr)
}
