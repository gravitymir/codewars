package main

import (
	"fmt"
)

//https://www.codewars.com/kata/525f50e3b73515a6db000b83
//Write a function that accepts an array
//of 10 integers (between 0 and 9),
//that returns a string of those numbers in the form of a phone number.

func main() {
	arr := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	resultString := CreatePhoneNumber(arr)
	fmt.Println(resultString)
}

//CreatePhoneNumber is returns "(123) 456-7890"
func CreatePhoneNumber(d [10]uint) (phone string) {
	country := fmt.Sprintf("%v%v%v", d[0], d[1], d[2])
	index := fmt.Sprintf("%v%v%v", d[3], d[4], d[5])
	number := fmt.Sprintf("%v%v%v%v", d[6], d[7], d[8], d[9])
	return fmt.Sprintf("(%v) %v-%v ", country, index, number)
}
