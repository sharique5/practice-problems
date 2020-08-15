package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func productExceptSelf(nums []int) []int {
	var res []int
	var r2l []int
	var l2r []int
	// init all arrays
	for i := 0; i < len(nums); i++ {
		res = append(res, 0)
		l2r = append(l2r, 0)
		r2l = append(r2l, 0)
	}
	// l2r
	temp := 1
	for i := 0; i < len(nums); i++ {
		temp *= nums[i]
		l2r[i] = temp

	}
	// r2l
	temp = 1
	for i := len(nums) - 1; i >= 0; i-- {
		temp *= nums[i]
		r2l[i] = temp
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			res[i] = r2l[1]
		} else if i == len(nums)-1 {
			res[i] = l2r[i-1]
		} else {
			res[i] = l2r[i-1] * r2l[i+1]
		}
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), ",")
	var nums []int
	for i := range line {
		num, err := strconv.Atoi(line[i])
		if err == nil {
			nums = append(nums, num)
		}
	}
	fmt.Println(productExceptSelf(nums))
}
