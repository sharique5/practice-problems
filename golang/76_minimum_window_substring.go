package main

import (
	"fmt"
	"strings"
)

func isTinS(freqS map[string]int, freqT map[string]int) bool {
	for key, tCount := range freqT {
		if sCount, ok := freqS[key]; ok {
			if tCount > sCount {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	if len(t) == 1 {
		if strings.Index(s, t) > -1 {
			return t
		}
		return ""
	}
	freqT := make(map[string]int)
	for i := 0; i < len(t); i++ {
		tempCh := string(t[i])
		if _, ok := freqT[tempCh]; ok {
			freqT[tempCh]++
		} else {
			freqT[tempCh] = 1
		}
	}
	freqS := make(map[string]int)
	l, r := 0, 1
	freqS[string(s[l])] = 1
	tempStr := string(s[r])
	if _, ok := freqS[tempStr]; ok {
		freqS[tempStr]++
	} else {
		freqS[tempStr] = 1
	}
	var res string = s
	var windowFound bool = false
	for l <= r && r < len(s) {
		if isTinS(freqS, freqT) {
			windowFound = true
			temp := s[l : r+1]
			if len(temp) < len(res) {
				res = temp
			}
			freqS[string(s[l])]--
			l++
		} else {
			r++
			if r < len(s) {
				tempStr = string(s[r])
				if _, ok := freqS[tempStr]; ok {
					freqS[tempStr]++
				} else {
					freqS[tempStr] = 1
				}
			}
		}
	}
	if windowFound == true {
		return res
	}
	return ""
}

func main() {
	var inps, inpt string
	fmt.Scanln(&inps)
	fmt.Scanln(&inpt)
	fmt.Println(inps, inpt)
	fmt.Println(minWindow(inps, inpt))
}
