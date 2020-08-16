package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		st, end := i+1, n-1
		for st < end {
			if nums[i]+nums[st]+nums[end] == 0 {
				temp := []int{nums[i], nums[st], nums[end]}
				result = append(result, temp)
				st += 1
				for st < end && nums[st] == nums[st-1] {
					st += 1
				}
				end -= 1
				for st < end && nums[end] == nums[end+1] {
					end -= 1
				}
			} else if nums[i]+nums[st]+nums[end] < 0 {
				st += 1
				for st < end && nums[st] == nums[st-1] {
					st += 1
				}
			} else {
				end -= 1
				for st < end && nums[end] == nums[end+1] {
					end -= 1
				}
			}
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), ", ")
	var nums []int
	for i := 0; i < len(line); i++ {
		num, err := strconv.Atoi(line[i])
		if err == nil {
			nums = append(nums, num)
		}
	}
	fmt.Println(threeSum(nums))
}
