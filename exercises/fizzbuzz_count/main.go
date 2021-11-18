package main

import "fmt"

// Problem statement: Given a string S, count the number of instances of "fizz", "buzz", and "fizzbuzz" in S

func main() {

	strs := []string{
		"fizz buzz fizzbuzz",                      // 2, 2, 1
		"fizzfizz fizz buzz fizzbuz",              // 4, 1, 0
		"fizzfizzbuzz fizz buzz fizzbuzz",         // 4, 3, 2
		"fiz fizz buzzbuzz fizzbuzz buzz fizzfiz", // 3, 4, 1
	}

	for _, str := range strs {
		counts := Count(str)
		fmt.Printf("'%s'\n", str)
		for _, count := range counts {
			fmt.Printf("  %s = %d\n", count.Word, count.Count)
		}
		fmt.Printf("\n")
	}

}

func Count(str string) []*WordCount {
	counts := []*WordCount{
		{"fizz", 0, 0},
		{"buzz", 0, 0},
		{"fizzbuzz", 0, 0},
	}

	for i := 0; i < len(str); i++ {
		for j := range counts {
			if counts[j].Word[counts[j].CurrIndex] == str[i] {
				if counts[j].CurrIndex == len(counts[j].Word)-1 {
					counts[j].Count++
					counts[j].CurrIndex = 0
				} else {
					counts[j].CurrIndex++
				}
			} else {
				counts[j].CurrIndex = 0
				if counts[j].Word[counts[j].CurrIndex] == str[i] {
					counts[j].CurrIndex++
				}
			}
		}
	}
	return counts
}

type WordCount struct {
	Word      string
	Count     int
	CurrIndex int
}
