package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxProfit(prices []int) int {
	ans, l, r := 0, 0, 0
	for r < len(prices) {
		if prices[l] <= prices[r] {
			diff := prices[r] - prices[l]
			if diff > ans {
				ans = diff
			}
			r++
		} else {
			l++
		}
	}
	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	splittedInput := strings.Split(line, ",")
	var nums []int
	for i := range splittedInput {
		num, err := strconv.Atoi(splittedInput[i])
		if err == nil {
			nums = append(nums, num)
		}
	}
	fmt.Println(maxProfit(nums))
}
