package main

import "fmt"

// 構造体
type person struct {
	name   string
	age    int
	smokes bool
	friend *person
}

func main() {
	// インスタンスを生成
	myster := person{}

	fmt.Println("name:", myster.name)
	fmt.Println("age:", myster.age)
	fmt.Println("smokes:", myster.smokes)
	fmt.Println("friend:", myster.friend)
}
