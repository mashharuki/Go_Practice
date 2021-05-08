package main

import "fmt"

func main() {
	intvals3 := []int{98, 125, 232, 147, 486}
	intvals3[1] -= 100

	fmt.Printf("このスライスが見ている配列の要素数は%d¥n", len(intvals3))
	fmt.Printf("最後の要素は%d¥n", intvals3[len(intvals3)-1])
	fmt.Printf("1番目の要素は%d¥n", intvals3[1])
}
