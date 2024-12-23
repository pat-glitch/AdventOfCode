package main

import (
	"bufio"
	"fmt"
	"os"
)

// Region struct holds information about area and perimeter of a region
type Region struct {
	Area      int
	Perimeter int
}

// Helper function to check if a cell is within bounds and matches the region's character
func isValid(x, y int, grid [][]rune, visited [][]bool, char rune) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && !visited[x][y] && grid[x][y] == char
}

// Perform a flood fill to calculate area and perimeter of a region
func floodFill(x, y int, grid [][]rune, visited [][]bool, char rune) Region {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // Up, Down, Left, Right
	queue := [][]int{{x, y}}
	visited[x][y] = true
	area, perimeter := 0, 0

	for len(queue) > 0 {
		cx, cy := queue[0][0], queue[0][1]
		queue = queue[1:]
		area++

		// Check all 4 directions
		for _, dir := range directions {
			nx, ny := cx+dir[0], cy+dir[1]
			if isValid(nx, ny, grid, visited, char) {
				visited[nx][ny] = true
				queue = append(queue, []int{nx, ny})
			} else if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[0]) || grid[nx][ny] != char {
				// Increment perimeter if the neighboring cell is outside or a different type
				perimeter++
			}
		}
	}

	return Region{Area: area, Perimeter: perimeter}
}

func calculateCost(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	totalCost := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] {
				region := floodFill(i, j, grid, visited, grid[i][j])
				cost := region.Area * region.Perimeter
				totalCost += cost
			}
		}
	}

	return totalCost
}

func main() {
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	totalCost := calculateCost(grid)
	fmt.Println("Total cost of fencing all regions:", totalCost)
}
