package main

import (
	"fmt"
	"strings"
)

//https://www.codewars.com/kata/5390bac347d09b7da40006f6/train/go
//Not Jaden-Cased: "How can mirrors be real if our eyes aren't real"
//Jaden-Cased:     "How Can Mirrors Be Real If Our Eyes Aren't Real"
func main() {
	str := "how  can mirrors be real if our eyes aren't real"

	resultString := ToJadenCase(str)
	fmt.Println(resultString)
}

func ToJadenCase(str string) string {

	arr := []rune(str)
	new := make([]rune, len(str))
	for k, v := range arr {
		new[k] = v
		if k > 0 && arr[k-1] <= 64 && arr[k] >= 65 && arr[k] <= 122 {
			f := strings.Title(string(arr[k]))
			new[k] = []rune(f)[0]
		} else if k == 0 {
			f := strings.Title(string(arr[0]))
			new[k] = []rune(f)[0]
		}
	}

	return string(new)
}
