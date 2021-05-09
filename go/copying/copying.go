package main

import "fmt"

func main() {
	inta := 5
	intb := inta
	inta = 7
	fmt.Printf("intaは%d¥n", inta)
	fmt.Printf("intbは%d¥n", intb)

	intb = 9
	fmt.Printf("intaは%d¥n", inta)
	fmt.Printf("intbは%d¥n", intb)

	fmt.Printf("intaのアドレスは %p¥n", &inta)
	fmt.Printf("intbのアドレスは %p¥n", &intb)
}
