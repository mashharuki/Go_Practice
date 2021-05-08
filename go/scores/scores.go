package main

import "fmt"

func main() {
	scores := [][3]int{
		{50, 55, 48}, {72, 85, 66}, {32, 100, 54}, {67, 68, 56}, {87, 67, 73}}

	for i := 0; i < len(scores); i++ {
		sum := 0
		for k := 0; k < 3; k++ {
			sum += scores[i][k]
		}
		fmt.Printf("受験者%d：平均値%d点¥n", i+1, sum/3)
	}
}
