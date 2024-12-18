package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const GridSize = 71 // 0-70 for both X and Y
const MaxBytes = 1024

// Direction vectors for moving up, down, left, or right
var directions = [][2]int{
	{-1, 0}, // Up
	{1, 0},  // Down
	{0, -1}, // Left
	{0, 1},  // Right
}

func readInput(filename string) ([][2]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var bytes [][2]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")
		if len(coords) != 2 {
			continue
		}
		x, err1 := strconv.Atoi(coords[0])
		y, err2 := strconv.Atoi(coords[1])
		if err1 != nil || err2 != nil {
			continue
		}
		bytes = append(bytes, [2]int{x, y})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return bytes, nil
}

func bfs(grid [][]bool) int {
	// Queue for BFS: each entry contains (x, y, steps)
	type state struct {
		x, y, steps int
	}
	queue := []state{{0, 0, 0}}

	// Visited array to avoid revisiting
	visited := make([][]bool, GridSize)
	for i := range visited {
		visited[i] = make([]bool, GridSize)
	}
	visited[0][0] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// If we reach the bottom-right corner, return the number of steps
		if current.x == GridSize-1 && current.y == GridSize-1 {
			return current.steps
		}

		// Explore neighbors
		for _, dir := range directions {
			nx, ny := current.x+dir[0], current.y+dir[1]

			// Check bounds and whether the cell is corrupted or visited
			if nx >= 0 && ny >= 0 && nx < GridSize && ny < GridSize && !grid[nx][ny] && !visited[nx][ny] {
				visited[nx][ny] = true
				queue = append(queue, state{nx, ny, current.steps + 1})
			}
		}
	}

	// If no path found, return -1
	return -1
}

func main() {
	// Read the input data
	bytes, err := readInput("inputdata.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Initialize the grid
	grid := make([][]bool, GridSize)
	for i := range grid {
		grid[i] = make([]bool, GridSize)
	}

	// Mark corrupted cells for the first MaxBytes entries
	for i := 0; i < MaxBytes && i < len(bytes); i++ {
		x, y := bytes[i][0], bytes[i][1]
		if x >= 0 && x < GridSize && y >= 0 && y < GridSize {
			grid[x][y] = true
		}
	}

	// Find the shortest path using BFS
	steps := bfs(grid)

	if steps == -1 {
		fmt.Println("No path to the exit.")
	} else {
		fmt.Printf("The minimum number of steps to reach the exit is: %d\n", steps)
	}
}
