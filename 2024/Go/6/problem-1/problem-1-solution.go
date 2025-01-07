package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MAX_ROWS = 1000
	MAX_COLS = 1000
)

type Position struct {
	x, y int
}

func main() {
	var grid [MAX_ROWS][MAX_COLS]rune
	var rows, cols int

	// Read the grid from the file
	readInputFile("inputdata.txt", &grid, &rows, &cols)

	// Calculate the distinct points visited
	distinctCount := markPath(&grid, rows, cols)
	fmt.Printf("Distinct points visited: %d\n", distinctCount)
}

func readInputFile(filename string, grid *[MAX_ROWS][MAX_COLS]rune, rows *int, cols *int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	*rows = 0
	for scanner.Scan() {
		line := scanner.Text()
		*cols = len(line)
		for i := 0; i < *cols; i++ {
			grid[*rows][i] = rune(line[i])
		}
		*rows++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
}

func markPath(grid *[MAX_ROWS][MAX_COLS]rune, rows, cols int) int {
	// Directions: Up, Right, Down, Left
	directions := [4]Position{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}
	direction := 0 // Start facing up
	distinctCount := 0

	// Find the starting position of '^'
	var current Position
	found := false
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '^' {
				current.x = i
				current.y = j
				grid[i][j] = 'X' // Mark the starting position
				distinctCount++
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		fmt.Println("Error: Starting position '^' not found.")
		os.Exit(1)
	}

	// Simulate the movement
	for {
		nextX := current.x + directions[direction].x
		nextY := current.y + directions[direction].y

		// Check if `^` exits the grid
		if nextX < 0 || nextX >= rows || nextY < 0 || nextY >= cols {
			break
		}

		// Check for an obstacle
		if grid[nextX][nextY] == '#' {
			direction = (direction + 1) % 4 // Rotate right
		} else {
			// Move to the next cell
			current.x = nextX
			current.y = nextY

			// Mark the cell if not already visited
			if grid[nextX][nextY] != 'X' {
				grid[nextX][nextY] = 'X'
				distinctCount++
			}
		}
	}

	return distinctCount
}
