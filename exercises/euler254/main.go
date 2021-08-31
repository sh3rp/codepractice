package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func sg(intStr string) int {
	return 0
}

func g(intStr string) int {
	return 0
}

func sf(intStr string) int {
	ints, _ := strToIntArray(intStr)

	total := 0

	for _, num := range ints {
		total += num
	}

	return total
}

func f(integers []int) int {
	total := 0

	for _, eachNum := range integers {
		factorial := 0
		for i := 1; i <= eachNum; i++ {
			factorial = factorial * i
		}
		total += factorial
	}

	return total
}

func intArrayToStr(integers []int) string {
	intStr := ""
	for _, num := range integers {
		el := strconv.Itoa(num)
		intStr = intStr + el
	}
	return intStr
}

func strToIntArray(str string) ([]int, error) {
	ints := []int{}

	for i := 0; i < len(str); i++ {
		integer, err := strconv.Atoi(string(str[i]))
		if err != nil {
			return nil, fmt.Errorf("error at index %d: not an int", i)
		}
		ints = append(ints, integer)
	}

	return ints, nil
}
