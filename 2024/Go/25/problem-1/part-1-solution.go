package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Schematic represents either a lock or key pattern
type Schematic struct {
	heights []int
	isLock  bool
}

// parseSchematic converts a group of lines into a single Schematic
func parseSchematic(lines []string) Schematic {
	if len(lines) == 0 {
		return Schematic{}
	}

	// Determine if this is a lock (# at top) or key (# at bottom)
	isLock := strings.Contains(lines[0], "#")
	width := len(lines[0])
	heights := make([]int, width)

	// For each column
	for col := 0; col < width; col++ {
		count := 0
		if isLock {
			// Count from top down
			for row := 0; row < len(lines); row++ {
				if lines[row][col] == '#' {
					count++
				} else {
					break
				}
			}
		} else {
			// Count from bottom up
			for row := len(lines) - 1; row >= 0; row-- {
				if lines[row][col] == '#' {
					count++
				} else {
					break
				}
			}
		}
		heights[col] = count
	}

	return Schematic{
		heights: heights,
		isLock:  isLock,
	}
}

// canFitTogether checks if a lock and key can fit together
func canFitTogether(lock, key Schematic, totalHeight int) bool {
	if len(lock.heights) != len(key.heights) {
		return false
	}

	for i := 0; i < len(lock.heights); i++ {
		if lock.heights[i]+key.heights[i] > totalHeight {
			return false
		}
	}
	return true
}

func main() {
	// Open input file
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentLines []string
	var schematics []Schematic

	// Read all schematics
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" && len(currentLines) > 0 {
			// Process completed schematic
			schematics = append(schematics, parseSchematic(currentLines))
			currentLines = nil
		} else if line != "" {
			currentLines = append(currentLines, line)
		}
	}

	// Don't forget to process the last schematic
	if len(currentLines) > 0 {
		schematics = append(schematics, parseSchematic(currentLines))
	}

	// Separate locks and keys
	var locks []Schematic
	var keys []Schematic
	for _, s := range schematics {
		if s.isLock {
			locks = append(locks, s)
		} else {
			keys = append(keys, s)
		}
	}

	// Count fitting pairs
	totalHeight := 7 // Based on example input
	fitCount := 0
	for _, lock := range locks {
		for _, key := range keys {
			if canFitTogether(lock, key, totalHeight) {
				fitCount++
			}
		}
	}

	fmt.Printf("Number of unique lock/key pairs that fit: %d\n", fitCount)
}
