package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"io/ioutil"
	"os"
)

func main() {
	// 画像ファイルを読み込む
	reader, ferr := os.Open("wow.png")

	if ferr != nil {
		fmt.Println("ファイルの読み込みエラーです。", ferr)
		os.Exit(1)
	}

	// デコードして情報を取り出す。
	img, name, derr := image.Decode(reader)

	if derr != nil {
		fmt.Println("画像の変換エラーです。", derr)
		os.Exit(1)
	} else {
		fmt.Println(name, "形式のデータを取得しました。")
	}
	// ファイルを閉じる。
	defer reader.Close()

	marks := []string{"*", "+", "-"}
	var marksStr string

	// 画質一つ一つについて、グレースケールで数値を取得する。
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			// グレー情報に変換して取得する。
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			grayness := c.Y / (255 / 3)

			if c.Y == 255 {
				grayness = 2
			}
			marksStr += marks[grayness]
		}
		marksStr += "¥n"
	}

	// バイトデータに変換する。
	wdata := []byte(marksStr)
	// ファイルに書き出す。
	werr := ioutil.WriteFile("ascil_art.txt", wdata, 0777)

	if werr != nil {
		fmt.Println("ファイルの書き込みエラーです。", werr)
		os.Exit(1)
	} else {
		fmt.Println("ファイルを保存しました")
	}
}
