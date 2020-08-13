package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			return true
		} else {
			m[nums[i]] = nums[i]
		}
	}
	return false
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
	fmt.Println(containsDuplicate(nums))
}
