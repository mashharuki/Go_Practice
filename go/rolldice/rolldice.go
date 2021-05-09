package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 構造体
type dice struct {
	val int
}

func rollDice(d dice, c chan int) {
	// ランダム
	rand.Seed(time.Now().UnixNano())

	for {
		// 10ミリ秒ごとに振る
		time.Sleep(10 * time.Millisecond)
		v := rand.Intn(10)
		if d.val == v {
			fmt.Printf("出ました！¥n")
			break
		} else {
			fmt.Printf("%dか...%dではないな¥n", v, d.val)
		}
	}
	c <- d.val
}

func main() {
	// インスタンスを生成
	d1 := dice{2}
	d2 := dice{6}
	// チャンネルを作る
	c := make(chan int)
	// rollDice関数を呼び出す。
	go rollDice(d1, c)
	go rollDice(d2, c)

	x, y := <-c, <-c

	fmt.Printf("%dが出ました！¥n", x)
	fmt.Printf("%dが出ました！¥n", y)
}
