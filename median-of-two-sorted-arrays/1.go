package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	index1 := 0
	index2 := 0
	var median int

    lenSum := len1+len2
    if lenSum == 0 {
    	return 0
	}

	remainder := lenSum%2

	if remainder == 1 {
		median = (lenSum)/2

		if len1 == 0 {
			return float64(nums2[median])
		}
		if len2 == 0 {
			return float64(nums1[median])
		}
		index := -1
		value := 0
		for index < median {
			index ++

			if index1 >= len1 {
				value = nums2[index2]
				index2 ++
				continue
			}

			if index2 >= len2 {
				value = nums1[index1]
				index1 ++
				continue
			}

			if nums1[index1] <= nums2[index2] {
				value = nums1[index1]
				index1 ++
			} else {
				value = nums2[index2]
				index2 ++
			}
		}
		return float64(value)

	} else {
		medianLeft := (len1+len2)/2 -1
		medianRight := (len1+len2)/2

		if len1 == 0 {
			return float64(nums2[medianLeft] + nums2[medianRight]) / 2
		}
		if len2 == 0 {
			return float64(nums1[medianLeft] + nums1[medianRight]) / 2
		}

		index := -1
		value := 0
		value2 := 0
		for index < medianRight {
			fmt.Println(index, index1, index2)
			index ++
			value2 = value

			if index1 >= len1 {
				value = nums2[index2]
				index2 ++
				continue
			}

			if index2 >= len2 {
				value = nums1[index1]
				index1 ++
				continue
			}

			if nums1[index1] <= nums2[index2] {
				value = nums1[index1]
				index1 ++
			} else {
				value = nums2[index2]
				index2 ++
			}

		}
		return float64(value + value2) / 2
	}
	return float64(nums1[index1] + nums2[index2]) /2
}

func main() {
	// nums1 	:= []int{1, 2}
	// nums2 := []int{3, 4}
	//
	nums1 	:= []int{100001}
	nums2 := []int{100000, 100001}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
