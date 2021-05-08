package main

import "fmt"

// インターフェース
type animaldata interface {
	// dedscribeメソッドを定義する。
	describe() string
}

// 構造体
type spec struct {
	name  string
	group string
}

type dog struct {
	sp     spec
	weight int
}

type cat struct {
	sp    spec
	btype string
}

func (sp spec) identify() string {
	return sp.group + "の" + sp.name + "さん"
}

// 犬の種類判定する。
func (d dog) classify() string {
	// 体重により分岐
	if d.weight > 25 {
		return "大型犬"
	}
	if d.weight < 10 {
		return "小型犬"
	}
	return "中型犬"
}

func (d dog) describe() string {
	// 関数を呼び出す。
	dscr := d.sp.identify() + ","
	dscr += "は" + d.classify()
	return dscr
}

// 猫のインスタンスを生成
func newcat(name string, group string, typenum int) cat {
	// specインスタンスを生成する。
	sp := spec{name, group}
	// チェックする。
	if typenum < 0 || typenum > 5 {
		typenum = 5
	}
	// タイプを設定する。
	btypes := [6]string{"オリエンタル", "コピー", "セミコピー", "フォーリン", "セミフォーリン", "サブスタンシャル"}
	return cat{sp, btypes[typenum]}
}

func (c cat) describe() string {
	// 関数を呼び出す。
	dscr := c.sp.identify() + ","
	dscr += "は" + c.btype
	return dscr
}

// メイン関数
func main() {
	// インターフェース実装
	animals := []animaldata{
		dog{spec{"コケ", "コーギー"}, 15},
		newcat("クロ", "雑種", 2),
		newcat("プリンス", "ペルシャ", 1),
		dog{spec{"ルドルフ", "ハスキー"}, 27},
	}
	// メソッドを呼び出す。
	for i := 0; i < len(animals); i++ {
		fmt.Println(animals[i].describe())
	}
}
