package main

import (
	"fmt"
	"math"
)

func characterReplacement(s string, k int) int {
	var res int = 0
	var freq []int
	for i := 0; i < 26; i++ {
		freq = append(freq, 0)
	}
	l, frequent := 0, 0
	for i := 0; i < len(s); i++ {
		freq[s[i]-'A']++
		frequent = int(math.Max(float64(frequent), float64(freq[s[i]-'A'])))
		for (i - l - frequent + 1) > k {
			freq[s[l]-'A']--
			l++
		}
		res = int(math.Max(float64(res), float64(i-l+1)))
	}
	return res
}

func main() {
	var inp string
	var k int
	fmt.Scanln(&inp)
	fmt.Scanf("%d", &k)
	fmt.Println(characterReplacement(inp, k))
}
