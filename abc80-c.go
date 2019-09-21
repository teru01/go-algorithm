package main

import (
	"fmt"
	"math"
)

var shops [][]int
var ben [][]int
var n int
var ans float64

func dfs(ary []int, d int) {
	if d == 10 {
		// fmt.Printf("%v ", ary)
		var q int
		for q=0; q<10; q++ {
			if ary[q] != 0 {
				break;
			}
		}
		if q == 10 {
			return
		}
		benefit := 0
		// o_c := make([]int, 0)
		for i := 0; i < n; i++ {
			c := 0
			for j := 0; j < 10; j++ {
				if ary[j]&shops[i][j] == 1 {
					// fmt.Println("hit: %v", j)
					c++
				}
			}
			benefit += ben[i][c]
			// o_c = append(o_c, ben[i][c])
		}
		if ans <= float64(benefit) {
			// if benefit == 0 {
			// 	fmt.Println(ary)
			// 	fmt.Println(ben)
			// }
			// fmt.Println(ary)
			ans = float64(benefit)
		}
		// fmt.Printf("%v\n", benefit)
		return
	}
	ary = append(ary, 0)
	dfs(ary, d+1)
	ary = ary[:len(ary)-1]
	ary = append(ary, 1)
	dfs(ary, d+1)
}

func main() {
	ans = -math.Pow10(8)
	fmt.Scan(&n)
	shops = make([][]int, n)
	for i := 0; i < n; i++ {
		shops[i] = make([]int, 10)
		for j:=0; j<10; j++ {
			fmt.Scan(&shops[i][j])
		}
	}

	ben = make([][]int, n)
	for i := 0; i < n; i++ {
		ben[i] = make([]int, 11)
		for j:=0; j<11; j++ {
			fmt.Scan(&ben[i][j])
		}
	}
	ary := make([]int, 0)
	// for i:=0; i<n; i++ {
	// 	for j:=0; j<10; j++ {
	// 		fmt.Printf("%v ", shops[i][j])
	// 	}
	// 	fmt.Print("\n")
	// }
	dfs(ary, 0)
	fmt.Println(ans)
}
