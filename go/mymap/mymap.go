package main

import "fmt"

// 構造体
type guest struct {
	name string
	age  int
	job  string
}

func (g guest) present() string {
	return fmt.Sprintf("%sさん、%d歳、職業：%s", g.name, g.age, g.job)
}

func main() {
	// マップを定義する
	legs := map[string]int{
		"bird":       2,
		"cat":        4,
		"grasshoper": 6,
		"octopus":    8,
		"squid":      10,
	}

	// マップの内容を出力する。
	for k, v := range legs {
		fmt.Printf("A %s has %s legs¥n", k, v)
	}

	// 空のマップを作成する
	rooms := make(map[int]guest)
	// マップにキーと値(この場合は構造体)を入れる。
	rooms[101] = guest{"山田", 29, "会社員"}
	rooms[208] = guest{"川崎", 21, "学生"}
	rooms[312] = guest{"高崎", 45, "公務員"}

	for k, v := range rooms {
		fmt.Printf("%d号室: ", k)
		fmt.Println(v.present())
	}
}
