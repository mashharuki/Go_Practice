package main

import "fmt"

func add(a int, b int) (result int, state string) {
	result = a + b
	state = "a + b"
	return
}

func sub(a int, b int) (result int, state string) {
	result = a - b
	state = "a - b"
	return
}

func mul(a int, b int) (result int, state string) {
	result = a * b
	state = "a * b"
	return
}

func div(a int, b int) (result int, state string) {
	result = a / b
	state = "a / b"
	return
}

func main() {
	// 関数のデータ型の要素を持つ配列を用意する。(要素数：4つ)
	oprs := []func(int, int) (int, string){
		add, sub, mul, div,
	}

	for i := 0; i < len(oprs); i++ {
		result, state := oprs[i](18, 6)
		fmt.Printf("%s = %d¥n", state, result)
	}
}
