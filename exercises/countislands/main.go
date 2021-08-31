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

func isSafe(M [][]int, row, col int, visited [][]bool) bool {
	return (row >= 0) && (row < len(M)) && (col >= 0) && (col < len(M[0])) && M[row][col] == 1 && !visited[row][col]
}

func DFS(M [][]int, row, col int, visited [][]bool) {
	rowNbr := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	colNbr := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	visited[row][col] = true

	for k := 0; k < 8; k++ {
		if isSafe(M, row+rowNbr[k], col+colNbr[k], visited) {
			DFS(M, row+rowNbr[k], col+colNbr[k], visited)
		}
	}
}

func countIslands(M [][]int) int {
	visited := make([][]bool, len(M))
	for i := range visited {
		visited[i] = make([]bool, len(M[0]))
	}

	count := 0

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
