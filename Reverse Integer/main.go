package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseNumber(12345678909))
}

func reverseNumber(num int) (res int) {
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return
}

// func threeSum(nums []int) (three [][]int) {
// 	var remainder []int

// 	for k, v := range nums {
// 		remainder = append(remainder, v)
// 		if (k+1)%3 == 0 {
// 			three = append(three, remainder)
// 			remainder = []int{}
// 		}
// 	}
// 	return three
// }
