package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// fmt.Println(nums, len(nums))
	dp := nums[0]
	ans := dp
	for i := 1; i < len(nums); i++ {
		if dp > 0 {
			dp = nums[i] + dp
		} else {
			dp = nums[i]
		}
		if dp > ans {
			ans = dp
		}
		// fmt.Println(nums[i], ans, dp)
	}
	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	splittedInput := strings.Split(line, " ")
	var nums []int
	for i := range splittedInput {
		num, err := strconv.Atoi(splittedInput[i])
		if err == nil {
			nums = append(nums, num)
		}
	}
	fmt.Println(maxSubArray(nums))
}
