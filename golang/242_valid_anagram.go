package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	sVal := strings.Split(s, "")
	sort.Strings(sVal)
	tVal := strings.Split(t, "")
	sort.Strings(tVal)
	return strings.Join(sVal, "") == strings.Join(tVal, "")
}

func main() {
	var s string
	fmt.Scanln(&s)
	var t string
	fmt.Scanln(&t)
	fmt.Println(isAnagram(s, t))
}
