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

// Collect unique trails using DFS
func dfsCollectTrails(x, y, height int, trails map[string]bool, currentPath [][]int) {
	// Check if already visited on this path
	if visited[x][y] {
		return
	}

	// Create a copy of the current trail and add current position
	newPath := make([][]int, len(currentPath)+1)
	copy(newPath, currentPath)
	newPath[len(currentPath)] = []int{x, y}

	// Check if this is a destination (9)
	if grid[x][y] == 9 {
		// Create a unique trail signature
		pathSig := generatePathSignature(newPath)
		trails[pathSig] = true
	}

	visited[x][y] = true

	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]

		if isValidMove(nx, ny, height) {
			dfsCollectTrails(nx, ny, height+1, trails, newPath)
		}
	}

	visited[x][y] = false // Backtrack
}

// Generate a unique signature for a path
func generatePathSignature(path [][]int) string {
	// Create a signature that captures the entire unique path
	sig := ""
	for _, coord := range path {
		sig += fmt.Sprintf("%d,%d|", coord[0], coord[1])
	}
	return sig
}

// Calculate sum of trailhead ratings
func calculateTrailheadRatings() int {
	totalRatings := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 { // Trailhead
				// Reset visited matrix
				for r := 0; r < rows; r++ {
					for c := 0; c < cols; c++ {
						visited[r][c] = false
					}
				}

				// Track unique trails
				trails := make(map[string]bool)

				// Start DFS with initial empty trail
				dfsCollectTrails(i, j, 0, trails, [][]int{})

				// Calculate rating (total number of unique trails)
				trailRating := len(trails)

				totalRatings += trailRating
				fmt.Printf("Trailhead at (%d, %d) has rating: %d\n", i, j, trailRating)
			}
		}
	}

	return totalRatings
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

	// Calculate and print total trailhead ratings
	totalTrailheadRatings := calculateTrailheadRatings()
	fmt.Println("Total Trailhead Ratings:", totalTrailheadRatings)
}
