package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Keypad layout for the numeric keypad
var numericKeypad = [][]string{
	{"7", "8", "9"},
	{"4", "5", "6"},
	{"1", "2", "3"},
	{"", "0", "A"},
}

// Directions for movement
var directions = map[string][2]int{
	"^": {-1, 0}, // Up
	"v": {1, 0},  // Down
	"<": {0, -1}, // Left
	">": {0, 1},  // Right
}

// Validates if a position is within the bounds of the numeric keypad
func isValidPosition(x, y int) bool {
	if x < 0 || x >= len(numericKeypad) || y < 0 || y >= len(numericKeypad[0]) {
		return false
	}
	return numericKeypad[x][y] != ""
}

// Finds the shortest path to type a number sequence on the numeric keypad
func findShortestPath(code string) int {
	startX, startY := 3, 2 // Start at the "A" key
	sequenceLength := 0

	for _, digit := range code {
		target := string(digit)
		steps, endX, endY := bfs(startX, startY, target)
		sequenceLength += steps
		startX, startY = endX, endY
	}

	return sequenceLength
}

// BFS to find shortest path to target key
func bfs(startX, startY int, target string) (int, int, int) {
	type state struct {
		x, y, steps int
	}
	queue := []state{{startX, startY, 0}}
	visited := make(map[[2]int]bool)
	visited[[2]int{startX, startY}] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if numericKeypad[current.x][current.y] == target {
			return current.steps, current.x, current.y
		}

		for dir, delta := range directions {
			nextX, nextY := current.x+delta[0], current.y+delta[1]
			if isValidPosition(nextX, nextY) && !visited[[2]int{nextX, nextY}] {
				visited[[2]int{nextX, nextY}] = true
				queue = append(queue, state{nextX, nextY, current.steps + 1})
			}
		}
	}

	return 0, startX, startY // Should never reach here
}

// Calculate complexity for a single code
func calculateComplexity(code string, sequenceLength int) int {
	numericValue, _ := strconv.Atoi(strings.TrimLeft(code, "0"))
	return sequenceLength * numericValue
}

func main() {
	// Hardcoded input data
	codes := []string{
		"129A",
		"974A",
		"805A",
		"671A",
		"386A",
	}

	totalComplexity := 0
	for _, code := range codes {
		sequenceLength := findShortestPath(code)
		totalComplexity += calculateComplexity(code, sequenceLength)
	}

	fmt.Println("Total Complexity:", totalComplexity)
}
