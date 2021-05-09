package main

import "fmt"

// 構造体
type roll struct {
	round int
	score int
}

// @return アドレス
func begin() *roll {
	fmt.Println("サイコロを投げます")
	return &roll{0, 0}
}

// サイコロを振る関数
func throw(r *roll, x int) {
	r.round++
	// 偶数なら加算する。
	if x%2 == 0 {
		r.score++
	}

	fmt.Printf("%d回目： スコア=%d¥n", r.round, r.score)
}

func main() {
	// インスタンス生成(アドレスを返す。)
	myroll := begin()
	// 関数を実行する。 (&が不要)
	throw(myroll, 6)
	throw(myroll, 5)
	throw(myroll, 2)
	throw(myroll, 4)
	throw(myroll, 3)
}
