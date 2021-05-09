package main

import "fmt"

// 構造体
type element struct {
	value string
	next  *element
}

type llist struct {
	head *element
	tail *element
}

// 追加するための関数
func (list *llist) add(newvalue string) {
	// インスタンスを生成
	elm := element{newvalue, nil}
	// 先頭かどうかチェック
	if list.head == nil {
		list.head = &elm
	} else {
		elm.isnextto(list.tail)
	}
	list.tail = &elm
}

// 次のポインタに設定する
func (elm *element) isnextto(prevelm *element) {
	prevelm.next = elm
}

// 値を取得する関数
func (elm *element) getValue() string {
	return elm.value
}

// リストの中身を全て表示する関数
func (list *llist) listall() {
	elm := list.head

	for {
		fmt.Println(elm.getValue())
		if elm.next == nil {
			break
		}
		elm = elm.next
	}
}

// メイン関数
func main() {
	// リストインスタンスを生成する。
	mylist := llist{nil, nil}
	mylist.add("クマ")
	mylist.add("マンドリル")
	mylist.add("ルリビタキ")
	mylist.add("キリン")
	mylist.add("あっ")

	mylist.listall()
}
