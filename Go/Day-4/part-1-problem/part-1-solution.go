package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("inputpuzzle.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read file into a grid
	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Dimensions of the grid
	rows := len(grid)
	if rows == 0 {
		fmt.Println("The file is empty or invalid.")
		return
	}
	cols := len(grid[0])

	// Word to search
	word := "XMAS"
	wordLen := len(word)

	// Function to check if the word exists in a given direction
	countWord := func(row, col, dRow, dCol int) int {
		for i := 0; i < wordLen; i++ {
			r := row + i*dRow
			c := col + i*dCol
			if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != word[i] {
				return 0
			}
		}
		return 1
	}

	// Directions: right, down, diagonal down-right, diagonal down-left,
	// reversed versions of these
	directions := [][2]int{
		{0, 1},   // Right
		{1, 0},   // Down
		{1, 1},   // Diagonal down-right
		{1, -1},  // Diagonal down-left
		{0, -1},  // Left (reversed right)
		{-1, 0},  // Up (reversed down)
		{-1, -1}, // Diagonal up-left (reversed down-right)
		{-1, 1},  // Diagonal up-right (reversed down-left)
	}

	// Count occurrences
	totalCount := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				totalCount += countWord(r, c, dir[0], dir[1])
			}
		}
	}

	fmt.Printf("The word '%s' appears %d times in the grid.\n", word, totalCount)
}
