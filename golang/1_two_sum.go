package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twoSum(nums []int, target int) []int {
	var emptyArr []int
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if index, ok := m[diff]; ok {
			res := []int{index, i}
			return res
		} else {
			m[nums[i]] = i
		}
	}
	return emptyArr
}

func main() {
	var nums []int
	var target int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	splittedInput := strings.Split(line, ", ")
	for i := range splittedInput {
		num, err := strconv.Atoi(splittedInput[i])
		if err == nil {
			nums = append(nums, num)
		}
	}
	scanner.Scan()
	target, err := strconv.Atoi(scanner.Text())
	if err != nil {
		target = -1
	}
	fmt.Println(twoSum(nums, target))
}
