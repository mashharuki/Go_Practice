package main

import "fmt"

func counter(adctr *int) {
	*adctr++
}

func freecounter(adctr *int, add int) {
	*adctr += add
}

func main() {
	// 変数を用意
	ctr := 0
	fctr := 10

	for i := 0; i < 5; i++ {
		// counter関数を呼び出す
		counter(&ctr)
		fmt.Printf("カウンタの値は%d¥n", ctr)

		freecounter(&fctr, ctr)
		fmt.Printf("フリーカウンタの値は%d¥n", fctr)
	}
}
