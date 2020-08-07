package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var res int = 0
	var st string = ""
	for _, c := range s {
		if strings.Contains(st, string(c)) {
			ind := strings.Index(st, string(c))
			st = st[ind+1:]
		}
		st += string(c)
		if len(st) > res {
			res = len(st)
		}
	}
	return res
}

func main() {
	var inp string
	fmt.Scanln(&inp)
	fmt.Println(lengthOfLongestSubstring(inp))
}
