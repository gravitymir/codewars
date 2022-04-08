package main

import (
	"fmt"
)

//https://www.codewars.com/kata/54da5a58ea159efa38000836/train/go
//Find the odd int

func main() {
	arr := []int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}
	resultString := FindOdd(arr)
	fmt.Println(resultString)
}

func FindOdd(seq []int) (res int) {
	res = 0
	for _, v := range seq {
		fmt.Println(v)
		if v%2 != 0 && v >= res {
			res = v
		}
	}
	return 0
}
