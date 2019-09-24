//me
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, sumA, sumB int
	fmt.Scan(&n)
	var a []int
	for i := 1; i <= n; i++ {
		var ai int
		fmt.Scan(&ai)
		a = append(a, ai)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for j := 0; j <= n-1; j++ {
		if j%2 == 0 {
			sumA += a[j]
		} else {
			sumB += a[j]
		}
	}
	fmt.Println(sumA - sumB)
}

//other person
package main
 
import (
	"fmt"
	"sort"
)
 
func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	result := 0
	for i := 0; i < n; i++ {
		if i&1 == 0 {
			result += a[i]
 
		} else {
			result -= a[i]
		}
	}
	fmt.Println(result)
}