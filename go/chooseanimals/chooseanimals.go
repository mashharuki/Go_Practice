package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 送られてきたデータによって異なる処理を行う
func choose(cf, cd, quit chan int) {
	for {
		select {
		case count := <-cf: // 偶数
			fmt.Printf("カエルが%d匹¥n", count)
		case count := <-cd: // 奇数
			fmt.Printf("アヒルが%d羽¥n", count)
		case <-quit: // 終了信号
			fmt.Println("終了")
			return
		}
	}
}

func countRandom(cf, cd, quit chan int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		// ランダムに数を発生させる
		count := rand.Intn(10) + 1
		// 奇数か偶数かをチェックする
		if (count % 2) == 0 {
			cf <- count
		} else {
			cd <- count
		}
	}
	// 終了信号を送る。
	quit <- 1
}

func main() {
	// チャンネルを作成する。
	cf := make(chan int)
	cd := make(chan int)
	quit := make(chan int)
	// 関数を呼び出す。
	go countRandom(cf, cd, quit)
	// 選択する
	choose(cf, cd, quit)
}
