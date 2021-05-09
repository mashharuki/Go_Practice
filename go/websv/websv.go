package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// 要求が来たときに実行する
	http.Handle("/", http.HandlerFunc(makePage))
	http.ListenAndServe(":8782", nil)
}

// makePage関数
func makePage(resw http.ResponseWriter, req *http.Request) {
	// ファイルを読み込む
	content, err := ioutil.ReadFile("indexhtml.txt")

	if err != nil {
		fmt.Println("テンプレートファイル読み込み失敗", err)
		os.Exit(1)
	} else {
		templatestr := string(content)
		// 文字列からHTMLを作成する。
		templ := template.Must(template.New("sendname").Parse(templatestr))
		// 要求を受け取ってページを書き出す。
		templ.Execute(resw, req.FormValue("yourname"))
	}
}
