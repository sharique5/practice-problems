package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxArea(height []int) int {
	area, l, r := 0, 0, len(height)-1
	for l < r {
		temp := 0
		if height[l] < height[r] {
			temp = height[l] * (r - l)
		} else {
			temp = height[r] * (r - l)
		}

		if temp > area {
			area = temp
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return area
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
	fmt.Println(maxArea(nums))
}
