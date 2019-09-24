package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	count := 0
	for {
		for d := 0; d < n; d++ {
			if a[d]%2 != 0 {
				fmt.Println(count)
				return
			}
			a[d] /= 2
		}
		count++
	}
}
