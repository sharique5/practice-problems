package main

import (
	"fmt"
	"strings"
)

func isAlphaNumeric(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9')
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(s string) bool {
	var st string
	for _, c := range s {
		if isAlphaNumeric(c) {
			st += strings.ToLower(string(c))
		}
	}
	return st == reverse(st)
}

func main() {
	var s string
	fmt.Scanln(&s)
	fmt.Println(isPalindrome(s))
}
