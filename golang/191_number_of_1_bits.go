package main

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	var count int
	for num > 0 {
		// fmt.Println("it:: ", num)
		count += int(num & 1)
		num >>= 1
	}
	return count
}

func main() {
	var i uint32
	fmt.Scanf("%d", &i)
	fmt.Println(i)
	fmt.Println(hammingWeight(i))
}
