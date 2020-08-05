package main

import "fmt"

// IsPalin is to check for palindromic string
func IsPalin(s string) bool {
	n := len(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func countSubstrings(s string) int {
	var cnt int = 0
	var n int = len(s)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n+1; j++ {
			st := s[i:j]
			if IsPalin(st) {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	var inp string
	fmt.Scanln(&inp)
	fmt.Println(countSubstrings(inp))
}
