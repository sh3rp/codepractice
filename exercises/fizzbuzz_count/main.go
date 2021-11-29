package main

import "fmt"

// Problem statement: Given a string S, count the number of instances of "fizz", "buzz", and "fizzbuzz" in S

var searchTerms = []string{
	"fizz",
	"buzz",
	"fizzbuzz",
}

func main() {

	strs := []string{
		"fizz buzz fizzbuzz",                      // 2, 2, 1
		"fizzfizz fizz buzz fizzbuz",              // 4, 1, 0
		"fizzfizzbuzz fizz buzz fizzbuzz",         // 4, 3, 2
		"fiz fizz buzzbuzz fizzbuzz buzz fizzfiz", // 3, 4, 1
	}

	for _, str := range strs {
		counts := Count(searchTerms,str)
		fmt.Printf("%s = %d, %s = %d, %s = %d\n",searchTerms[0], counts[0], searchTerms[1], counts[1], searchTerms[2], counts[2])
	}

}

func Count(terms []string, str string) []int {
	states := make([]*TermState,len(terms))
	for idx, term := range terms {
		states[idx] = &TermState{term,0,0}
	}

	for i := 0; i < len(str); i++ {
		for j := range states {
			if states[j].Word[states[j].CurrIndex] == str[i] {
				if states[j].CurrIndex == len(states[j].Word)-1 {
					states[j].Count++
					states[j].CurrIndex = 0
				} else {
					states[j].CurrIndex++
				}
			} else {
				states[j].CurrIndex = 0
				if states[j].Word[states[j].CurrIndex] == str[i] {
					states[j].CurrIndex++
				}
			}
		}
	}

	counts := make([]int,len(states))
	for idx,state := range states {
		counts[idx] = state.Count
	}

	return counts
}

type TermState struct {
	Word      string
	Count     int
	CurrIndex int
}
