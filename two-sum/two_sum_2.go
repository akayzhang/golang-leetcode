package main

import (
	"fmt"
)

func twoSum2(nums []int, target int) []int {
	hash := make(map[int]int)

	for index, num := range nums {
		hash[num] = index
	}
	for index, num := range nums {
		if v, ok := hash[(target - num)]; ok {
			return []int{index, v}
		}
	}
	return []int{}
}

func main()  {
	nums := []int{2, 7, 11, 15}
	target := 9

	res := twoSum2(nums, target)
	fmt.Println(res)
}
