package main

import (
	"fmt"
)

//https://leetcode.com/problems/two-sum/
//Two Sum
func main() {
	sl := []int{
		2, 7, 11, 15,
	}
	fmt.Println(twoSum(sl, 9))
}
func twoSum(nums []int, target int) (sl []int) {
	for k := range nums {
		for k2 := range nums {
			if k == k2 {
				continue
			}
			if nums[k]+nums[k2] == target {
				return []int{k, k2}
			}
		}
	}
	return
}

/*
Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]

*/
