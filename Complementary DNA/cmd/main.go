package main

import (
	"fmt"
	"strings"
)

//https://www.codewars.com/kata/554e4a2f232cdd87d9000038/train/go
// "ATTGC" --> "TAACG"
// "GTAT" --> "CATA"
// dnaStrand []        `shouldBe` []
// dnaStrand [A,T,G,C] `shouldBe` [T,A,C,G]
// dnaStrand [G,T,A,T] `shouldBe` [C,A,T,A]
// dnaStrand [A,A,A,A] `shouldBe` [T,T,T,T]

func main() {
	dna := "TACG"

	resultString := DNAStrand(dna)
	fmt.Println(resultString)
}

func DNAStrand(dna string) string {
	return strings.NewReplacer(
		"A", "T",
		"T", "A",
		"G", "C",
		"C", "G").Replace(dna)
}
