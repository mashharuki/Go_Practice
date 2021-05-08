package main

import "fmt"

// 構造体
type patient struct {
	name   string
	height int
	weight int
	chest  int
	ubp    int
	hbp    int
	lbp    int
}

// 健康診断結果インスタンスの初期化
func newpatient(name string) patient {
	return patient{name, 0, 0, 0, 0, 0, 0}
}

// 身長・体重測定についての関数
func (pat patient) bodys(height int, weight int) patient {
	// 身長と体重をセットして返す。
	pat.height = height
	pat.weight = weight
	return pat
}

// 血圧を設定する関数
func (pat patient) bloodp(hbp int, lbp int) patient {
	// 血圧をセットして返す。
	pat.hbp = hbp
	pat.lbp = lbp
	return pat
}

// BMIを求めて測定する関数
func (pat patient) bmi() string {
	// 0以下でないことをチェックする
	if pat.weight*pat.height <= 0 {
		return "測定エラー"
	}
	// BMIの算出
	bmi := (pat.weight * 10000) / (pat.height * pat.height)
	// BMIの値に応じて処理を分岐する。
	if bmi < 19 {
		return "もっと食べましょう"
	} else if bmi > 25 {
		return "運動しましょう"
	}
	return "バランスの取れた体格です。"
}

// 血圧の範囲をチェックする。
func (pat patient) bprange() string {
	// 0以下でないことをチェックする
	if pat.hbp*pat.lbp <= 0 {
		return "測定エラー"
	}
	// 血圧をチェックする
	if pat.hbp < 100 {
		return "血圧が低いかもしれません"
	}
	if pat.hbp < 140 && pat.lbp < 90 {
		return "血圧は正常です"
	}
	return "血圧に注意が必要です"
}

// 測定する。
func (pat patient) report() string {
	report := pat.name + "さん: ¥n"
	report += pat.bmi() + "¥n"
	report += pat.bprange()
	return report
}

// メインの関数
func main() {
	// インスタンス生成
	patients := []patient{
		newpatient("HM").bodys(170, 78).bloodp(120, 80),
		newpatient("SK").bodys(165, 88).bloodp(135, 92),
		newpatient("TN").bodys(156, 39).bloodp(122, 70),
		newpatient("AT").bodys(160, 55).bloodp(98, 66),
	}

	fmt.Println("***健康診断の結果です***")

	for i := 0; i < len(patients); i++ {
		fmt.Println(patients[i].report())
		fmt.Println()
	}
	fmt.Println("***なお、胸囲・座高測定は省力します。***")
}
