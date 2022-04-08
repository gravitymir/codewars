package main

import (
	"fmt"
)

//https://www.codewars.com/kata/55fd2d567d94ac3bc9000064/train/go
//Given the triangle of consecutive odd numbers:
func main() {
	resultString := RowSumOddNumbers(101)
	fmt.Println(resultString)
	resultString = RowSumOddNumbers(2)
	fmt.Println(resultString)
}

func RowSumOddNumbers(n int) int {
	return n * n * n
}
