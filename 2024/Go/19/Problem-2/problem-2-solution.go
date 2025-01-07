package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to count the number of ways a design can be formed using the available patterns
func countDesignWays(patterns []string, design string) int {
	// Dynamic Programming to count the number of ways
	ways := make([]int, len(design)+1)
	ways[0] = 1 // Base case: There is one way to form an empty design

	for i := 1; i <= len(design); i++ {
		for _, pattern := range patterns {
			if len(pattern) <= i && ways[i-len(pattern)] > 0 {
				if design[i-len(pattern):i] == pattern {
					ways[i] += ways[i-len(pattern)]
				}
			}
		}
	}

	return ways[len(design)]
}

func main() {
	// Open the input file
	file, err := os.Open("inputdata.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read towel patterns
	scanner.Scan()
	patterns := strings.Split(scanner.Text(), ", ")

	// Skip the blank line
	scanner.Scan()

	designs := []string{}

	// Read designs
	for scanner.Scan() {
		design := scanner.Text()
		if design != "" {
			designs = append(designs, design)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Count the total number of ways for all designs
	totalWays := 0

	for _, design := range designs {
		ways := countDesignWays(patterns, design)
		fmt.Printf("Design: %s, Ways: %d\n", design, ways)
		totalWays += ways
	}

	fmt.Printf("Total number of ways: %d\n", totalWays)
}
