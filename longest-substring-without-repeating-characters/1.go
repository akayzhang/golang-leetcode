package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {

	if s == "" {
		return 0
	}


	len := len(s)

	hash := make(map[byte]int)

	low, fast := 0,0
	length  := 1
	for fast < len {
		if v,ok := hash[s[fast]];ok {
			for low <= v {
				delete(hash,s[low])
				low++
			}
		} else {
			hash[s[fast]] = fast
			if fast - low + 1 > length {
				length = fast - low + 1
			}
			fast ++
		}
	}
	return length
}

func main()  {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}



