package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	one := 0
	two := 1
	len := len(nums)
	for nums[one] + nums[two] != target {
		two++
		if two == len {
			one++
			two = one+1
		}
	}
	return []int{one, two}
}

func main()  {
	nums := []int{2, 7, 11, 15}
	target := 9

	res := twoSum(nums, target)
	fmt.Println(res)
}