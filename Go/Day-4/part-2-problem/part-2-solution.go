package main

import (
	"bufio"
	"fmt"
	"os"
)

// Function to load the grid from a file
func loadGrid(filename string) ([]string, int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, 0, err
	}

	if len(grid) == 0 {
		return nil, 0, 0, fmt.Errorf("empty grid")
	}

	rows := len(grid)
	cols := len(grid[0])
	return grid, rows, cols, nil
}

// Function to check if a cell is within the grid bounds
func isValid(r, c, rows, cols int) bool {
	return r >= 0 && r < rows && c >= 0 && c < cols
}

// Function to check if an X-MAS pattern exists centered at (r, c)
func isXMas(grid []string, r, c, rows, cols int) bool {
	// Ensure the center is 'A'
	if grid[r][c] != 'A' {
		return false
	}

	// Define the diagonal positions for the X pattern
	topLeft := [2]int{r - 1, c - 1}
	topRight := [2]int{r - 1, c + 1}
	bottomLeft := [2]int{r + 1, c - 1}
	bottomRight := [2]int{r + 1, c + 1}

	// Validate all positions are within bounds
	if !isValid(topLeft[0], topLeft[1], rows, cols) ||
		!isValid(topRight[0], topRight[1], rows, cols) ||
		!isValid(bottomLeft[0], bottomLeft[1], rows, cols) ||
		!isValid(bottomRight[0], bottomRight[1], rows, cols) {
		return false
	}

	// Check the two diagonals for "MAS" or "SAM"
	diag1 := string([]byte{grid[topLeft[0]][topLeft[1]], grid[r][c], grid[bottomRight[0]][bottomRight[1]]})
	diag2 := string([]byte{grid[topRight[0]][topRight[1]], grid[r][c], grid[bottomLeft[0]][bottomLeft[1]]})

	// Valid X-MAS patterns
	return (diag1 == "MAS" || diag1 == "SAM") && (diag2 == "MAS" || diag2 == "SAM")
}

// Function to count all X-MAS patterns in the grid
func countXMas(grid []string, rows, cols int) int {
	count := 0

	// Iterate over all possible centers of the X
	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if isXMas(grid, r, c, rows, cols) {
				count++
			}
		}
	}

	return count
}

func main() {
	// File containing the grid
	filename := "inputpuzzle.txt"

	// Load the grid
	grid, rows, cols, err := loadGrid(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Count X-MAS patterns
	count := countXMas(grid, rows, cols)

	// Output the result
	fmt.Printf("Total X-MAS patterns found: %d\n", count)
}
