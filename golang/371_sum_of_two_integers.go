package main

import "fmt"

func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	for b != 0 {
		c := a & b
		a = a ^ b
		b = c << 1
	}

	return a
}

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	// fmt.Println("value::: ", a, b)
	fmt.Println(getSum(a, b))
}
