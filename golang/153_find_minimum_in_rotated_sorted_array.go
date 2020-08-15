package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findMin(nums []int) int {
	l := 0
	h := len(nums) - 1
	if h < 0 {
		return 0
	}
	for (l + 1) < h {
		m := l + (h-l)/2
		if nums[m] > nums[l] && nums[m] > nums[h] {
			l = m
		} else {
			h = m
		}
	}
	if nums[l] < nums[h] {
		return nums[l]
	}
	return nums[h]
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
	fmt.Println(findMin(nums))
}
