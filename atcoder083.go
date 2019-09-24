package main

import "fmt"

func main() {
	var n, a, b, sum2 int
	fmt.Scan(&n, &a, &b)

	for i := 1; i <= n; i++ {
		sum1 := 0
		for j := i; j != 0; j /= 10 {
			sum1 += j % 10
		}
		if a <= sum1 && sum1 <= b {
			sum2 += i
		}

	}
	fmt.Println(sum2)
}
