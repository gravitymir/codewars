package main

import (
	"fmt"
	"strings"
)

//https://www.codewars.com/kata/5667e8f4e3f572a8f2000039/train/go
func main() {
	fmt.Println(Accum("ZpglnRxqenU"))
}

func getSub(length int, letter string) (s string) {
	s += strings.Title(letter)
	s += strings.Repeat(strings.ToLower(letter), length)
	return
}

func Accum(s string) string {
	arr := strings.Split(s, "")
	var resArr string

	for k := range arr {
		resArr += getSub(k, arr[k])
		if k < len(arr)-1 {
			resArr += "-"
		}
	}
	return resArr
}
