package main

import (
	"fmt"
	"sort"
)

//https://leetcode.com/problems/two-sum/
//Two Sum
func main() {
	sl := []int{
		2, 7, 11, 15,
	}
	twoSum(sl, 9)
}
func twoSum(nums []int, target int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return i < j
	})
	fmt.Println(nums)
	return []int{1, 2}
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
