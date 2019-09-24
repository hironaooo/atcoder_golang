package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	m := make(map[int]bool)
	uniq := []int{}
	for _, num := range a {
		if !m[num] {
			m[num] = true
			uniq = append(uniq, num)
		}
	}
	fmt.Println(len(uniq))
}

// 他の人
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	m := make(map[int]int)
// 	var d int
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&d)
// 		m[d] = 0
// 	}
// 	fmt.Println(len(m))
// }
