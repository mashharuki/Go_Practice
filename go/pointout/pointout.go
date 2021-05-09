package main

import "fmt"

func main() {
	inta := 5
	adinta := &inta
	bdinta := &inta
	*bdinta = 9

	fmt.Printf("adintaの値は %p¥n", adinta)
	fmt.Printf("intaの値は %d¥n", inta)
	fmt.Printf("adintaの内容は %d¥n", *adinta)
	fmt.Printf("adintaのアドレスは %p¥n", &adinta)
	fmt.Printf("bdintaのアドレスは %p¥n", &bdinta)
}
