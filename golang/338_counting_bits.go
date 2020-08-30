package main

import "fmt"

func countBits(num int) []int {
	var ans []int
	for i := 0; i < num+1; i++ {
		temp := i
		cnt := 0
		for temp > 0 {
			if temp&1 == 1 {
				cnt++
			}
			temp >>= 1
		}
		ans = append(ans, cnt)
	}
	return ans
}

func main() {
	var inp int
	fmt.Scanf("%d", &inp)
	fmt.Println(countBits(inp))
}
