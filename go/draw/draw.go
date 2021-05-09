package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// 範囲内にあるかどうかチェックする関数
func inrange(x, a, b int) bool {
	if a < b && a <= x && x < b {
		return true
	}
	return false
}

// メイン関数
func main() {
	// 画像の縦横情報を取得する。
	const width = 350
	const height = 350
	const celllen = 50
	// イメージ図情報
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	// 各セルのcolorsの要素を割り当てる
	cells := [7][7]int{
		{0, 1, 0, 1, 0, 1, 0},
		{0, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 0},
		{0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 2, 0, 0, 0},
		{0, 2, 0, 2, 0, 2, 0},
		{0, 0, 2, 2, 2, 0, 0},
	}
	// 色を格納するための配列
	colors := [3]color.RGBA{
		color.RGBA{0xff, 0xff, 0xff, 0xff},
		color.RGBA{0xff, 0, 0, 0xff},
		color.RGBA{0, 0xff, 0, 0xff},
	}

	// 各セルごとに色を設定する
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// 縦横をそれぞれ7つに分割する。
			for ny := 0; ny < 7; ny++ {
				for nx := 0; nx < 7; nx++ {
					// 範囲内にあることをチェックする
					if inrange(x, nx*celllen, (nx+1)*celllen) && inrange(y, ny*celllen, (ny+1)*celllen) {
						img.Set(x, y, colors[cells[ny][nx]])
					}
				}
			}
		}
	}

	// 画像ファイルの作成
	f, ferr := os.Create("flower.png")
	if ferr != nil {
		fmt.Println("ファイル作成のエラー", ferr)
		os.Exit(1)
	}
	// 画像データを書き込む
	derr := png.Encode(f, img)

	if derr != nil {
		f.Close()
		fmt.Println("データ変換のエラー：", derr)
		os.Exit(1)
	}
	// ファイルを閉じる。
	ferr = f.Close()
	if ferr != nil {
		fmt.Println("ファイル終了のエラー：", ferr)
		os.Exit(1)
	}
}
