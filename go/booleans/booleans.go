package main

import "fmt"

func main() {
	a := true
	fmt.Printf("aの値は%t¥n", a)
	b := false
	fmt.Printf("a==bの値は%t¥n", a == b)
	c := 2
	d := 5
	fmt.Printf("c!=dの値は%t¥n", c != d)
	fmt.Printf("c<dの値は%t¥n", c < d)
	fmt.Printf("(-c)<(-d)の値は%t¥n", -1*c < -1*d)

	over1under3 := 1 < c && c < 3
	fmt.Printf("over1under3の値は%t¥n", over1under3)
	under5just5 := d < 5 || d == 5
	fmt.Printf("under5just5の値は%t¥n", under5just5)
	ctimesdnot10 := (c*d != 10)
	fmt.Printf("ctimesdnot10の値は%t¥n", ctimesdnot10)
}
