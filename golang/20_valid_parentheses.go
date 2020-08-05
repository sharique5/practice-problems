package main

import "fmt"

// Contains given string in array or not
func Contains(a [3]string, val string) bool {
	for _, c := range a {
		if c == val {
			return true
		}
	}
	return false
}

func isValid(s string) bool {
	var stack []string
	openingBraces := [3]string{"(", "{", "["}
	closingBraces := [3]string{")", "}", "]"}
	for _, c := range s {
		strC := string(c)
		if Contains(openingBraces, strC) {
			stack = append(stack, strC)
		} else if Contains(closingBraces, strC) {
			n := len(stack) - 1
			if n == -1 {
				return false
			}
			topCh := stack[n]
			if (topCh == "(" && strC == ")") || (topCh == "{" && strC == "}") || (topCh == "[" && strC == "]") {
				stack = stack[:n]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	var inp string
	fmt.Scanln(&inp)
	fmt.Println(isValid(inp))
}
