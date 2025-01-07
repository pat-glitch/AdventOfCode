package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	maxRows = 200
	maxCols = 200
)

var (
	rows, cols int
	grid       [maxRows][maxCols]int
	visited    [maxRows][maxCols]bool
	dx         = []int{-1, 1, 0, 0}
	dy         = []int{0, 0, -1, 1}
)

// Check if a move is valid
func isValidMove(x, y, prevHeight int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols &&
		!visited[x][y] && grid[x][y] == prevHeight+1
}

// Collect reachable 9s using DFS
func dfsCollect9s(x, y, height int, reachable9s map[int]map[int]bool) {
	if reachable9s[x] == nil {
		reachable9s[x] = make(map[int]bool)
	}

	if grid[x][y] == 9 {
		reachable9s[x][y] = true
	}

	visited[x][y] = true

	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]

		if isValidMove(nx, ny, height) {
			dfsCollect9s(nx, ny, height+1, reachable9s)
		}
	}

	visited[x][y] = false // Backtrack
}

// Calculate total score with distinct 9s
func calculateTotalScoreDistinct() int {
	totalScore := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 { // Trailhead
				// Reset visited matrix
				for r := 0; r < rows; r++ {
					for c := 0; c < cols; c++ {
						visited[r][c] = false
					}
				}

				// Track distinct reachable 9s
				reachable9s := make(map[int]map[int]bool)

				dfsCollect9s(i, j, 0, reachable9s)

				// Count distinct 9s
				distinctCount := 0
				for _, row := range reachable9s {
					distinctCount += len(row)
				}

				totalScore += distinctCount
			}
		}
	}

	return totalScore
}

func main() {
	// Open input file
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read input map
	scanner := bufio.NewScanner(file)
	rows = 0

	for scanner.Scan() && rows < maxRows {
		line := scanner.Text()
		cols = len(line)

		for j, char := range line {
			digit, _ := strconv.Atoi(string(char))
			grid[rows][j] = digit
		}

		rows++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Calculate and print total score
	finalTotalScoreDistinct := calculateTotalScoreDistinct()
	fmt.Println(finalTotalScoreDistinct)
}
