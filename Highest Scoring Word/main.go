package main

import (
	"fmt"
	"strings"
)

//https://www.codewars.com/kata/57eb8fcdf670e99d9b000272/train/go
func main() {
	fmt.Println(High("man i need a taxi up to ubud"))
	fmt.Println(High("what time are we climbing up the volcano"))
	fmt.Println(High("aa b")) //aa
	fmt.Println(High("b aa")) //b
}

func High(s string) string {
	arr := strings.Fields(s)
	max, ind := 0, 0
	for i, v := range arr {
		sum := 0
		for k := range v {
			sum += int(v[k]) - 96
		}
		fmt.Println("sum:", sum, "word:", v)
		if max < sum {
			max = sum
			ind = i
			fmt.Println("sum:", sum, "word:", v)
		}
	}
	return arr[ind]
}
