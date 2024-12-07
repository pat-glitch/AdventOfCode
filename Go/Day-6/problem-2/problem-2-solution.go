package main

import (
	"bufio"
	"fmt"
	"os"
)

const SIZE = 130

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

var directions = [4][2]int{
	{-1, 0}, // NORTH
	{0, 1},  // EAST
	{1, 0},  // SOUTH
	{0, -1}, // WEST
}

func main() {
	// Open the input file
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var grid [SIZE][SIZE]rune
	var guardRow, guardCol int

	// Read the input file and fill the grid
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		for col, ch := range line {
			grid[row][col] = ch
			if ch == '^' {
				guardRow, guardCol = row, col
				grid[row][col] = '.' // Replace '^' with '.'
			}
		}
		row++
	}

	// Solve Part 1: Count visited cells
	part1 := simulateGuard(guardRow, guardCol, NORTH, &grid, false)

	// Solve Part 2: Count loop-causing obstacle placements
	part2 := countLoopCausingObstacles(guardRow, guardCol, &grid)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

// simulateGuard simulates the guard's movement and returns the count of visited cells (or detects loops)
func simulateGuard(guardRow, guardCol, direction int, grid *[SIZE][SIZE]rune, checkLoop bool) int {
	visited := make(map[[3]int]bool) // Tracks visited positions with direction
	visitedCells := make(map[[2]int]bool)
	count := 0

	for {
		state := [3]int{guardRow, guardCol, direction}
		if checkLoop && visited[state] {
			return 1 // Loop detected
		}
		visited[state] = true

		// Mark the cell as visited
		pos := [2]int{guardRow, guardCol}
		if !visitedCells[pos] {
			visitedCells[pos] = true
			count++
		}

		// Calculate the next position
		nextRow := guardRow + directions[direction][0]
		nextCol := guardCol + directions[direction][1]

		// Check if the next position is within bounds
		if nextRow < 0 || nextRow >= SIZE || nextCol < 0 || nextCol >= SIZE || grid[nextRow][nextCol] == '#' {
			// Turn right if out of bounds or an obstacle
			direction = (direction + 1) % 4
		} else {
			// Move to the next position
			guardRow, guardCol = nextRow, nextCol
		}
	}
}

// countLoopCausingObstacles tests all empty positions for loop-causing behavior
func countLoopCausingObstacles(startRow, startCol int, grid *[SIZE][SIZE]rune) int {
	count := 0

	for r := 0; r < SIZE; r++ {
		for c := 0; c < SIZE; c++ {
			if grid[r][c] == '.' {
				// Temporarily add an obstacle
				grid[r][c] = '#'

				// Check if it causes a loop
				if simulateGuard(startRow, startCol, NORTH, grid, true) == 1 {
					count++
				}

				// Restore the grid
				grid[r][c] = '.'
			}
		}
	}

	return count
}
