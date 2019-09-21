package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		n int
	)
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	ans := float64(math.Pow10(9))
	for _, v := range a {
		c := 0
		for v % 2 == 0 {
			v /= 2
			c++
		}
		ans = math.Min(float64(ans), float64(c))
	}
	fmt.Println(ans)
}
