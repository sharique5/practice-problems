package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) < 3 {
		for i := 0; i < len(nums); i++ {
			if nums[i] == target {
				return i
			}
		}
		return -1
	}

	left, right := 0, len(nums)-1
	if nums[left] == target {
		return left
	} else if nums[right] == target {
		return right
	}
	for true {
		if right-left <= 1 {
			return -1
		}
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] < target && target < nums[mid] {
			right = mid
		} else if nums[mid] < target && target < nums[right] {
			left = mid
		} else if nums[mid] > nums[left] {
			left = mid
		} else if nums[right] > nums[mid] {
			right = mid
		} else {
			return -1
		}
	}
	return -1
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
	var target int
	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())
	if err == nil {
		target = num
	}
	fmt.Scanf("%d", &target)
	fmt.Println(search(nums, target))
}
