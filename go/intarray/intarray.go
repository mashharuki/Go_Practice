package main

import "fmt"

func main() {
	intvals := [5]int{98, 125, 232, 147, 486}
	fmt.Printf("1番目の要素は%d¥n", intvals[1])
	fmt.Printf("0番目の要素は%d¥n", intvals[0])
	fmt.Printf("4番目の要素は%d¥n", intvals[4])
	fmt.Printf("2番目と3番目の要素を足すと%d¥n", intvals[2]+intvals[3])
	intvals[1] -= 100
	fmt.Printf("1番目の要素は%d¥n", intvals[1])
}
