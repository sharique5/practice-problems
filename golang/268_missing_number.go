package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func missingNumber(nums []int) int {
	sum := 0
	n := len(nums)
	for i := range nums {
		sum += nums[i]
	}
	total := (n * (n + 1)) / 2
	return total - sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), ",")
	var inp []int
	for i := 0; i < len(line); i++ {
		num, err := strconv.Atoi(line[i])
		if err == nil {
			inp = append(inp, num)
		}
	}
	fmt.Println(missingNumber(inp))
}
