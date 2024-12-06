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

var obstacleMap [SIZE][SIZE]int

// Directions for moving the guard
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

	var guardRow, guardCol, direction int = 0, 0, NORTH
	var mapGrid [SIZE][SIZE]rune
	var dirMap [SIZE][SIZE]int

	// Read the input file and fill the grid
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		for col, ch := range line {
			mapGrid[row][col] = ch
			if ch == '^' {
				guardRow, guardCol = row, col
				mapGrid[row][col] = '.' // replace ^ with . for easier processing
			}
		}
		row++
	}

	// Solve the problem
	count := solve(guardRow, guardCol, direction, 0, &mapGrid, &dirMap)

	// Print the final state of the grid and the result
	printMap(&mapGrid, &dirMap)
	fmt.Printf("Part 1 (Total Obstacle Count): %d\n", count)
}

// solve simulates the guard's movement and marks obstacles (if depth == 0) or detects loops (if depth == 1)
func solve(guardRow, guardCol, direction, depth int, mapGrid *[SIZE][SIZE]rune, dirMap *[SIZE][SIZE]int) int {
	count := 0

	for {
		// If we're searching for a loop and we've visited this cell with the same direction, return loop found
		if depth == 1 && mapGrid[guardRow][guardCol] == 'X' && (dirMap[guardRow][guardCol]&(1<<direction)) != 0 {
			return 1
		}

		// Mark the current position as visited ('X') and record the direction
		mapGrid[guardRow][guardCol] = 'X'
		dirMap[guardRow][guardCol] |= (1 << direction)

		// Calculate the next position based on the current direction
		nextRow, nextCol := guardRow+directions[direction][0], guardCol+directions[direction][1]

		// Check if the next position is within bounds
		if nextRow >= 0 && nextRow < SIZE && nextCol >= 0 && nextCol < SIZE {
			// If it's an obstacle, turn right (clockwise)
			if mapGrid[nextRow][nextCol] == '#' {
				direction = (direction + 1) % 4
			} else {
				// Otherwise, move the guard to the next cell
				guardRow, guardCol = nextRow, nextCol
			}

			// If we're checking for loop formation and there's no obstacle, add an obstacle and check for loops
			if depth == 0 && mapGrid[nextRow][nextCol] != 'X' && obstacleMap[nextRow][nextCol] == 0 {
				// Create a copy of the grid with the new obstacle
				newMap := *mapGrid
				newDirMap := *dirMap
				newMap[nextRow][nextCol] = '#'

				// Call solve again with the new obstacle
				obstacle := solve(guardRow, guardCol, (direction+1)%4, 1, &newMap, &newDirMap)
				if obstacle > 0 {
					obstacleMap[nextRow][nextCol] = 1
					count++
				}
				guardRow, guardCol = nextRow, nextCol
			}
		} else {
			// If the guard goes out of bounds, stop
			break
		}
	}

	// Return count if it's part of the loop detection (depth == 1)
	if depth == 1 {
		return 0
	}

	return count
}

// printMap prints the map with obstacles and direction indicators
func printMap(mapGrid *[SIZE][SIZE]rune, dirMap *[SIZE][SIZE]int) {
	for row := 0; row < SIZE; row++ {
		for col := 0; col < SIZE; col++ {
			if mapGrid[row][col] == 'X' {
				n, e, s, w := (dirMap[row][col]&(1<<NORTH)) != 0, (dirMap[row][col]&(1<<EAST)) != 0,
					(dirMap[row][col]&(1<<SOUTH)) != 0, (dirMap[row][col]&(1<<WEST)) != 0

				if (n || s) && (e || w) {
					fmt.Print("+")
				} else if n || s {
					fmt.Print("|")
				} else if e || w {
					fmt.Print("-")
				}
			} else {
				fmt.Print(string(mapGrid[row][col]))
			}
		}
		fmt.Println()
	}
}
