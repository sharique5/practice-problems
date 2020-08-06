package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	var st string
	st = "$#"
	for _, c := range s {
		st += string(c)
		st += "#"
	}
	st += "@"
	// fmt.Println("updated:: ", st)
	var p []int
	for i := 0; i < len(st); i++ {
		p = append(p, 0)
	}
	// fmt.Println("p::: ", p)
	c, r := 0, 0
	for i := 1; i < len(st)-1; i++ {
		if i < r {
			if r-i < p[2*c-i] {
				p[i] = r - i
			} else {
				p[i] = p[2*c-i]
			}
		}

		for st[i+(1+p[i])] == st[i-(1+p[i])] {
			p[i]++
		}

		if i+p[i] > r {
			c = i
			r = 1 + p[i]
		}
	}
	centerIndex, palLen := 0, 0
	for i := 0; i < len(p); i++ {
		if p[i] > palLen {
			palLen = p[i]
			centerIndex = i
		}
	}
	res := ""
	if st[centerIndex] != '#' && st[centerIndex] != '@' && st[centerIndex] != '$' {
		res = string(st[centerIndex])
	}
	for i := centerIndex + 1; i < centerIndex+palLen+1; i++ {
		if st[i] != '#' {
			res = string(st[i]) + res + string(st[i])
		}
	}
	return res
}

func main() {
	var inp string
	fmt.Scanln(&inp)
	fmt.Println(longestPalindrome(inp))
}
