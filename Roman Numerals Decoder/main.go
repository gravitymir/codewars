package main

import "fmt"

//https://www.codewars.com/kata/51b6249c4612257ac0000005/train/go
func main() {
	fmt.Println(Decode("MDCLXVI"))
	fmt.Println(Decode("IV"))

}

func Decode(roman string) (sum int) {
	var matrix = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	for k := range roman {
		if k+1 <= len(roman)-1 {
			if matrix[string(roman[k])] < matrix[string(roman[k+1])] {
				//fmt.Println("asddasda")
				sum -= matrix[string(roman[k])]
				continue
			}
		}
		sum += matrix[string(roman[k])]
	}
	return sum
}
