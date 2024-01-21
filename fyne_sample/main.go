package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

// メイン関数
func main() {
	a := app.New()

	w := a.NewWindow("Hello World!!")
	w.SetContent(
		widget.NewLabel("Hello World!!"),
	)

	w.ShowAndRun()
}
