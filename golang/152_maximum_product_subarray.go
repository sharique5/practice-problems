package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxProduct(nums []int) int {
	tMax := nums[0]
	tMin := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			temp := tMax
			tMax = tMin
			tMin = temp
		}
		if tMax*nums[i] > nums[i] {
			tMax = tMax * nums[i]
		} else {
			tMax = nums[i]
		}
		if tMin*nums[i] < nums[i] {
			tMin = tMin * nums[i]
		} else {
			tMin = nums[i]
		}
		if tMax > ans {
			ans = tMax
		}
	}
	return ans
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), ",")
	var nums []int
	for i := 0; i < len(line); i++ {
		num, err := strconv.Atoi(line[i])
		if err == nil {
			nums = append(nums, num)
		}
	}
	fmt.Println(maxProduct(nums))
}
