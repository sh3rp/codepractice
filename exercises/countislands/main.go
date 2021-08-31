package main

import "fmt"

// Exercise: Given a 2d integer array, where a 1 is "land" and a 0 is "water", find all the islands
// in the array

func main() {
	theMap := [][]int{
		{1, 1, 0, 0, 0},
		{0, 1, 0, 0, 1},
		{1, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1},
	}
	fmt.Printf("Number of islands: %d", countIslands(theMap))
}

// check for the following:
//  - are the row/col values within the boundaries of the matrix
//  - is the current element "land"
//  - has the current element been visited
//
//  if any are false, we should not continue the search
func isValid(M [][]int, row, col int, visited [][]bool) bool {
	return (row >= 0) && (row < len(M)) && (col >= 0) && (col < len(M[0])) && M[row][col] == 1 && !visited[row][col]
}

func DFS(M [][]int, row, col int, visited [][]bool) {
	// these are actually a set of offset coordinates around the current element
	// representing the surrounding elements we need to evaluation
	rowNbr := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	colNbr := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	// since we've evaluated this element, mark it as visited
	visited[row][col] = true

	// look around the element for any land
	for k := 0; k < 8; k++ {
		// did we find land?  if so, recurse and look for more
		if isValid(M, row+rowNbr[k], col+colNbr[k], visited) {
			DFS(M, row+rowNbr[k], col+colNbr[k], visited)
		}
	}
}

func countIslands(M [][]int) int {
	// our visited array used to track elements that have been visited (evaluated)
	visited := make([][]bool, len(M))
	for i := range visited {
		visited[i] = make([]bool, len(M[0]))
	}

	count := 0

	// start looking through the 2d matrix
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[0]); j++ {
			if M[i][j] == 1 && !visited[i][j] {
				DFS(M, i, j, visited)
				count++
			}
		}
	}

	return count
}
