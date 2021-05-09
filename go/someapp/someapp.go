package main

import "fmt"

// アプリケーションを想定した構造体
type someapp struct {
	// ユーザー名
	username string
	// 使用回数
	use int
	// 開いているかどうか
	isopen bool
}

// アプリケーションを開く関数
func (app *someapp) open() {
	// 開いているかチェック
	if app.isopen {
		fmt.Printf("%sさんのアプリはすでに開いています¥n¥n", app.username)
	} else {
		// 使用回数を加算する
		app.use++
		// 使用回数により処理を分岐する。
		if app.use > 2 {
			fmt.Printf("%sさんの試用期間は満了です。ご購入をご検討ください¥n¥n", app.username)
		} else {
			fmt.Printf("こんにちは%sさん、%d回目のご使用になります。¥n¥n", app.username, app.use)
			// フラグをONにする。
			app.isopen = true
		}
	}
}

// アプリケーションを閉じる関数
func (app *someapp) close() {
	// フラグの値をチェックする
	if app.isopen {
		fmt.Printf("さよなら%sさん、アプリを終了します¥n¥n", app.username)
		app.isopen = false
	}
}

// アプリケーションのインスタンスを生成する関数
// @return somepp 構造体
func newapp(username string) someapp {
	fmt.Printf("さよなら%sさん、初めてのご使用になります¥n¥n", username)
	return someapp{username, 1, true}
}

// メイン関数
func main() {
	fmt.Println("[ringoさんがアプリを使用開始]")
	rapp := newapp("ringo")
	rapp.close()

	fmt.Println("[ringoさんが再びアプリを使用]")
	rapp.open()
	rapp.close()

	fmt.Println("[mikanさんがアプリを使用開始]")
	mapp := newapp("mikan")
	mapp.close()

	fmt.Println("[mikanさんが再びアプリを使用]")
	mapp.open()

	fmt.Println("[ringoさんが再びアプリを使用]")
	rapp.open()

	fmt.Println("[mikanさんが再びアプリを使用]")
	mapp.open()
}
